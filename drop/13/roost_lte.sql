# INSERT INTO loottable_entries(probability, loottable_id, lootdrop_id, drop_limit, mindrop) VALUES (100, :todo_lt:, :todo_ld:, 1, 1);

INSERT INTO lootdrop_entries(lootdrop_id, item_id, chance) VALUES(:todo:, 13404, 100); # Bear Meat 92638 110223 23.47 a_vicious_mountain_bear (398004 lvl 56 chance 23.471%), a_mountain_bear (398005 lvl 52 chance 24.922%), a_mountain_bear_cub (398006 lvl 48 chance 22.673%), a_vicious_mountain_bear (398007 lvl 55 chance 23.471%), a_hungry_mountain_bear (398009 lvl 56 chance 22.269%), a_mountain_bear (398011 lvl 50 chance 24.922%), a_mountain_bear (398012 lvl 51 chance 24.922%), a_vicious_mountain_bear (398013 lvl 57 chance 23.471%), a_mountain_bear_cub (398016 lvl 49 chance 22.673%), a_mountain_bear_cub (398017 lvl 47 chance 22.673%)

INSERT INTO loottable_entries(probability, loottable_id, lootdrop_id, drop_limit, mindrop) VALUES 
(20, :roost_a_vicious_mountain_bear_lt:, :bear_meat_ld:),
(20, :roost_a_mountain_bear_lt:, :bear_meat_ld:),
(20, :roost_a_mountain_bear_cub_lt:, :bear_meat_ld:),
(20, :roost_a_hungry_mountain_bear_cub:, :bear_meat_ld:),