# Drop
drop is a project inside peq-expansions to reitemize all zones into smarter loot table and loot drops

## 1. NPC Dump

First step is to get a dump of NPCs known to spawn in a zone. (Note: this won't cover NPCs that are spawned via quests)
```sql
SELECT group_concat(id SEPARATOR ', '), NAME, level FROM npc_types WHERE id IN (SELECT npcid FROM spawnentry WHERE spawngroupid IN (SELECT spawngroupid FROM spawn2 WHERE zone = "blackburrow")) GROUP BY NAME;
```
Take results, and save as **zone**_npc.sql
Next, rename the data to this pattern: `UPDATE npc_types SET loottable_id = :a_burly_gnoll_lt: WHERE id IN (17003, 17024, 17033); # a_burly_gnoll 9`

## 2. Headhunter Dump

We need to double check existance of all noteworthy mobs inside peq's db, now. Modern EQ has the headhunter achievements to give us hints: e.g. thulehouse1 https://everquest.allakhazam.com/db/quest.html?quest=6860
We can take the listing resulted there, and dump it on the upper half of the **zone**_npc.sql file, and format it with - [ ] checkmarks to see which ones are missing.

Since we aren't focused on ADDING npcs, only modifying loot tables, we can go ahead and make placeholder tables for missing named NPC (if any) as new entries , e.g.:
```
# missing npc's but we want loot tables anyways for when they get added.
bonecracker_lt
darnor_the_terror_lord_lt
dreameater_lt
```

## 3. Loot Table Prep

Do a search for _lt inside **zone**_npc.sql, copy all entries resulted and paste to new file named **zone**_table.sql
format like: 
```sql
INSERT INTO loottable(name) VALUES ('alboct_vinn'); # :alboct_vinn_lt:
```

This preps each NPC to have a unique loot table, we can now freely link loot drop categories to this in the future.

## 4. Loot Drop Prep

Now it's time to get actual drops. This part going to be more situational, and one of the more time consuming steps.

Run this query:
```sql
SELECT lootdrop_entries.item_id,  items.name AS item_name, loottable_entries.loottable_id, lootdrop_entries.lootdrop_id, lootdrop_entries.chance, group_concat(concat("# ", npc_types.name, " (", npc_types.id, " lvl ", npc_types.level ,")") SEPARATOR ', ') AS npc
FROM lootdrop_entries INNER JOIN items ON items.id = lootdrop_entries.item_id 
INNER JOIN loottable_entries ON loottable_entries.lootdrop_id = lootdrop_entries.lootdrop_id 
INNER JOIN npc_types ON npc_types.loottable_id =  loottable_entries.loottable_id 
WHERE npc_types.loottable_id = loottable_entries.loottable_id AND lootdrop_entries.lootdrop_id IN(
    SELECT lootdrop_id FROM loottable_entries WHERE loottable_id IN (SELECT loottable_id FROM npc_types WHERE id IN (
        SELECT npcid FROM spawnentry WHERE spawngroupid IN (
            SELECT spawngroupid FROM spawn2 WHERE zone = "blackburrow")))
    ) GROUP BY item_name ORDER BY item_id; 
```

This is a of dump every item every npc drops in peq for a zone. Save this as **zone**_drop.sql

We want to go through each record now, and start structuring them to the pattern noted above.
example:
```
from: 
54445	Scuffed Weapon Crate of the Mercenary	95509	117758	0.24	# #Gristle (701086 lvl 88)
to (note the 5 is just placeholder):
(:todo:, 54445, 5), # Scuffed Weapon Crate of the Mercenary	95509	117758	0.24	# #Gristle (701086 lvl 88)
```

This list is ordered by item id's, so you can see hints on how items might be grouped. 
Make spaces between records and see if you can find patterns, rearranging each entry in what appears to be possible loot drop categories.

You'll notice lots of global items, too, you can either delete, or shift it to the bottom of the list to discard later..

Next, look for items that are dropped exclusively by a single npc. This is likely personal loot, and you can rename the :todo: fields on these records to that npc's name, e.g. :named_ld:.

Whenever you create a new _ld entry, you need to create it on top of the file. E.g., 
```sql
INSERT INTO lootdrop(name) VALUES ('name'); # :name_ld:
```

Next, let's name other loot drops based on their (likely) grouping scope, e.g., if a set of item drops appears to be gnoll casters carry it, you can replace the entries in that drop grouping from :todo: to :gnoll_caster_ld:, (don't forget to create a new lootdrop on top!)

If you see any out of era items, you can note it in the lootdrop's name, e.g. gnoll_caster_loy_ld. You can use this query to create the lootdrop:
```sql
INSERT INTO lootdrop(name, min_expansion) VALUES ('gnoll_caster_loy', 5); # :gnoll_cast_loy_ld:
```

## 5. Missing Loot Check

Alright! We're moving along. We got some rough groupings together, some obvious loot tables created, now let's double check itemization.

The most obvious NPCs to check against are the named mobs to start, so, grab a named NPC, and find them on allakhazam. You're trying to find out if any unique drops may be missing for the NPC.
For newer zones (HoT+), [eqresource](https://hot.eqresource.com/) is amazing for getting a break down of personalized loot. They have a [section](https://hot.eqresource.com/allnamehotlower.php) that breaks down named with all items of note they drop. They even include placeholder names and locations. (Tip: You can also use the items dropped and get a list of every relevant NPC, this is super helpful to knowing how to group even better)
