# based on https://almarsguides.com/eq/farming/loot/gems/
INSERT INTO lootdrop(name, min_expansion) VALUES ('global_jewelcrafting_gem_tss', -1); #:global_jewelcrafting_gem_tss_ld:

# level 50 is when they start dropping
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES(:global_jewelcrafting_gem_tss_ld:, 97001, 100); # Prestidigitase
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES(:global_jewelcrafting_gem_tss_ld:, 97002, 100); # Staurolite
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES(:global_jewelcrafting_gem_tss_ld:, 97003, 100); # Harmonagate
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES(:global_jewelcrafting_gem_tss_ld:, 97286, 100); # Taaffeite
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES(:global_jewelcrafting_gem_tss_ld:, 37856, 100); # Shaped Taaffeite
