INSERT INTO lootdrop(name) VALUES ('buckler'); # :buckler_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:buckler_ld:, 9001, 100); # Buckler

INSERT INTO lootdrop(name) VALUES ('kite_shield'); # :kite_shield_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:kite_shield_ld:, 9004, 100); # Kite Shield

INSERT INTO lootdrop(name) VALUES ('wooden_shield'); # :wooden_shield_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:wooden_shield_ld:, 9006, 100); # Wooden Shield

INSERT INTO lootdrop(name) VALUES ('small_buckler'); # :small_buckler_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:small_buckler_ld:, 9009, 100); # Small Buckler

INSERT INTO lootdrop(name) VALUES ('small_wooden_shield'); # :small_wooden_shield_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:small_wooden_shield_ld:, 9014, 100); # Small Wooden Shield

INSERT INTO lootdrop(name) VALUES ('round_shield'); # :round_shield_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:round_shield_ld:, 9002, 100); # Round Shield

INSERT INTO lootdrop(name) VALUES ('targ_shield'); # :targ_shield_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:targ_shield_ld:, 9003, 100); # Targ Shield

# dropped primarily by guards in small race towns
INSERT INTO lootdrop(name, min_expansion) VALUES ('small_round_shield', -1); #:small_round_shield_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:small_round_shield_ld:, 9010, 100); # Small Round Shield

INSERT INTO lootdrop(name, min_expansion) VALUES ('cracked_darkwood_shield', -1); #:cracked_darkwood_shield_ld:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:cracked_darkwood_shield_ld:, 9309, 100); # Cracked Darkwood Shield

INSERT INTO lootdrop(name) VALUES('qeynos_kite_shield'); #:qeynos_kite_shield:
INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES
(:qeynos_kite_shield_ld:, 9023, 100); # Qeynos Kite Shield