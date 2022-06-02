package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/projecteq/peq-expansions/db"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed:", err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := db.Init(ctx)
	if err != nil {
		return fmt.Errorf("db.Init: %w", err)
	}
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: drop <zone_shortname>")
	}

	shortName := os.Args[1]

	fmt.Println(shortName)
	zoneID, err := zoneID(ctx, shortName)
	if err != nil {
		return fmt.Errorf("zoneID: %w", err)
	}

	err = npcDump(ctx, zoneID, shortName)
	if err != nil {
		return fmt.Errorf("npcDump: %w", err)
	}

	err = dropDump(ctx, shortName)
	if err != nil {
		return fmt.Errorf("dropDump: %w", err)
	}
	lootdropDump(ctx, shortName)
	if err != nil {
		return fmt.Errorf("lootdropDump: %w", err)
	}
	loottableEntriesDump(ctx, shortName)
	if err != nil {
		return fmt.Errorf("lootTableEntriesDump: %w", err)
	}
	return nil
}

func zoneID(ctx context.Context, shortName string) (zoneIDNumber int, error error) {
	type dropRecord struct {
		Zoneidnumber int
	}

	query := `SELECT
	zoneidnumber
	FROM zone
	WHERE short_name = :short_name`
	args := map[string]interface{}{
		"short_name": shortName,
	}

	rows, err := db.Instance.NamedQueryContext(ctx, query, args)
	if err != nil {
		return 0, fmt.Errorf("query zone: %w", err)
	}

	for rows.Next() {
		drop := &dropRecord{}
		err = rows.StructScan(&drop)
		if err != nil {
			return 0, fmt.Errorf("query zone scan: %w", err)
		}
		zoneIDNumber = drop.Zoneidnumber
		return
	}
	return 0, fmt.Errorf("zoneIDNumber for %s not found", shortName)
}

func npcDump(ctx context.Context, zoneID int, shortName string) error {

	type npcRecord struct {
		Npcs  string
		Name  string
		Level int
	}

	path := fmt.Sprintf("%s_n.sql", shortName)

	minZone := zoneID*1000 - 1
	maxZone := zoneID*1000 + 999

	query := `SELECT
	group_concat(id SEPARATOR ', ') as npcs, name, level
	FROM npc_types
	WHERE id > :minZone AND id < :maxZone
	GROUP BY name`
	args := map[string]interface{}{
		"minZone": minZone,
		"maxZone": maxZone,
	}

	rows, err := db.Instance.NamedQueryContext(ctx, query, args)
	if err != nil {
		return fmt.Errorf("query npc_types: %w", err)
	}

	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	wt, err := os.Create(fmt.Sprintf("%s_lt.sql", shortName))
	if err != nil {
		return err
	}
	defer wt.Close()

	for rows.Next() {
		npc := &npcRecord{}
		err = rows.StructScan(&npc)
		if err != nil {
			return fmt.Errorf("query npc scan: %w", err)
		}

		npcTable := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(npc.Name), " ", "_"), "#", ""), "`", "")
		_, err = w.WriteString(fmt.Sprintf("UPDATE npc_types SET loottable_id = :%s_%s_lt: WHERE id in (%s); # %s %d\n", shortName, npcTable, npc.Npcs, npc.Name, npc.Level))
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}

		_, err = wt.WriteString(fmt.Sprintf("INSERT INTO loottable(name) VALUES('%s_%s'); # :%s_%s_lt:\n", shortName, npcTable, shortName, npcTable))
		if err != nil {
			return fmt.Errorf("writeTable: %w", err)
		}
	}
	return nil
}

func dropDump(ctx context.Context, shortName string) error {
	type lootRecord struct {
		Item_id      int
		Item_name    string
		Loottable_id int
		Lootdrop_id  int
		Chance       float32
		Npc          string
	}

	path := fmt.Sprintf("%s_lde.sql", shortName)

	query := `
	SELECT lootdrop_entries.item_id, items.name AS item_name, loottable_entries.loottable_id, lootdrop_entries.lootdrop_id, lootdrop_entries.chance, group_concat(concat(npc_types.name, " (", npc_types.id, " lvl ", npc_types.level ," chance ", lootdrop_entries.chance ,"%)") SEPARATOR ', ') AS npc
	FROM lootdrop_entries INNER JOIN items ON items.id = lootdrop_entries.item_id 
	INNER JOIN loottable_entries ON loottable_entries.lootdrop_id = lootdrop_entries.lootdrop_id 
	INNER JOIN npc_types ON npc_types.loottable_id =  loottable_entries.loottable_id 
	WHERE npc_types.loottable_id = loottable_entries.loottable_id AND lootdrop_entries.lootdrop_id IN(
		SELECT lootdrop_id FROM loottable_entries WHERE loottable_id IN (SELECT loottable_id FROM npc_types WHERE id IN (
			SELECT npcid FROM spawnentry WHERE spawngroupid IN (
				SELECT spawngroupid FROM spawn2 WHERE zone = :zone)))
		) GROUP BY item_name ORDER BY item_id;
`
	args := map[string]interface{}{
		"zone": shortName,
	}

	rows, err := db.Instance.NamedQueryContext(ctx, query, args)
	if err != nil {
		return fmt.Errorf("query loot: %w", err)
	}

	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	bosses := make(map[string][]string)
	trashes := []string{}

	for rows.Next() {
		loot := &lootRecord{}
		err = rows.StructScan(&loot)
		if err != nil {
			return fmt.Errorf("query loot scan: %w", err)
		}

		dump := fmt.Sprintf("INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES(:todo:, %d, %d); # %s %d %d %0.2f %s\n", loot.Item_id, 100, loot.Item_name, loot.Loottable_id, loot.Lootdrop_id, loot.Chance, loot.Npc)
		mobs := strings.Split(loot.Npc, ",")
		mobCount := 1
		lastMob := ""
		for _, mob := range mobs {
			if strings.Contains(mob, "(") {
				mob = strings.TrimSpace(mob[0:strings.Index(mob, "(")])
			}
			if lastMob == "" {
				lastMob = mob
			}
			if mob == lastMob {
				continue
			}
			mobCount++
			break
		}
		if mobCount > 1 {
			trashes = append(trashes, dump)
			continue
		}
		bosses[loot.Npc] = append(bosses[loot.Npc], dump)
	}

	for _, boss := range bosses {
		for _, dump := range boss {
			_, err = w.WriteString(dump)
			if err != nil {
				return fmt.Errorf("write boss dump: %s", err)
			}
		}
	}
	w.WriteString("\n")
	for _, dump := range trashes {
		_, err = w.WriteString(dump)
		if err != nil {
			return fmt.Errorf("write trash dump: %s", err)
		}
	}
	return nil
}

func lootdropDump(ctx context.Context, shortName string) error {
	w, err := os.Create(fmt.Sprintf("%s_ld.sql", shortName))
	if err != nil {
		return err
	}
	defer w.Close()
	w.WriteString("# INSERT INTO lootdrop(name) VALUES('todo'); # :todo_ld:\n")
	return nil
}

func loottableEntriesDump(ctx context.Context, shortName string) error {
	w, err := os.Create(fmt.Sprintf("%s_lte.sql", shortName))
	if err != nil {
		return err
	}
	defer w.Close()

	w.WriteString("# INSERT INTO loottable_entries(probability, loottable_id, lootdrop_id, drop_limit, mindrop) VALUES (100, :todo_lt:, :todo_ld:, 1, 1);\n")
	return nil
}
