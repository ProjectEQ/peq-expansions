INSERT INTO lootdrop(name, min_expansion) VALUES ('mammoth_meat', -1); #:mammoth_meat_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:mammoth_meat_ld:, 13085, 70), # Chunk of Meat
(:mammoth_meat_ld:, 13404, 30), # Mammoth Meat

INSERT INTO lootdrop(name, min_expansion) VALUES ('small_mammoth_tusks', -1); #:small_mammoth_tusks_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:small_mammoth_tusks_ld:, 10125, 100); # Small Mammoth Tusks
