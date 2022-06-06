INSERT INTO lootdrop(name, min_expansion) VALUES ('lightstone', -1); #:lightstone_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:lightstone_ld:, 10300, 100), # Lightstone
(:lightstone_ld:, 10400, 100), # Greater Lightstone
(:lightstone_ld:, 10299, 100); # Burned Out Lightstone

INSERT INTO lootdrop(name, min_expansion) VALUES ('wisp_essence', -1); #:wisp_essence_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:wisp_essence_ld:, 27399, 100); # Wisp Essence


INSERT INTO lootdrop(name, min_expansion) VALUES ('enchanted_wisp_globe', -1); #:enchanted_wisp_globe_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:enchanted_wisp_globe_ld:, 27415, 100); # Enchanted Wisp Globe

INSERT INTO lootdrop(name, min_expansion) VALUES ('vial_of_smoke', -1); #:vial_of_smoke_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:vial_of_smoke_ld:, 9923, 100); # Vial of Smoke

INSERT INTO lootdrop(name, min_expansion) VALUES ('third_riddle_for_the_troll', -1); #:third_riddle_for_the_troll_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:third_riddle_for_the_troll_ld:, 18654, 100); # Third Riddle for the Troll
