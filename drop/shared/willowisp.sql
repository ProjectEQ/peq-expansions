INSERT INTO lootdrop(name, min_expansion) VALUES ('lightstone', -1); #:lightstone_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:lightstone_ld:, 10300, 100), # Lightstone
(:lightstone_ld:, 10400, 100), # Greater Lightstone
(:lightstone_ld:, 10299, 100); # Burned Out Lightstone

INSERT INTO lootdrop(name, min_expansion) VALUES ('wisp_essence', -1); #:wisp_essence_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:wisp_essence_ld:, 27399, 100); # Wisp Essence
