#--
UPDATE npc_types SET loottable_id = :a_burly_gnoll_lt: WHERE id IN (17003, 17024, 17033); # a_burly_gnoll 9
UPDATE npc_types SET loottable_id = :giant_snake_lt: WHERE id IN (17013, 17017, 17018); # a_giant_snake 7
UPDATE npc_types SET loottable_id = :a_gnoll_brewer_lt: WHERE id IN (17036); # a_gnoll_brewer 11
UPDATE npc_types SET loottable_id = :a_gnoll_commander_lt: WHERE id IN (17021, 17037); # a_gnoll_commander 14
UPDATE npc_types SET loottable_id = :a_gnoll_courier_lt: WHERE id IN (17030); # a_gnoll_courier 30
UPDATE npc_types SET loottable_id = :a_gnoll_guardsman_lt: WHERE id IN (17009, 17010, 17012, 17016, 17031, 17039); # a_gnoll_guardsman 9
UPDATE npc_types SET loottable_id = :a_gnoll_lt: WHERE id IN (17002, 17008, 17019); # a_gnoll 6
UPDATE npc_types SET loottable_id = :a_gnoll_pup_lt: WHERE id IN (17014); # a_gnoll_pup 1
UPDATE npc_types SET loottable_id = :a_gnoll_tactician_lt: WHERE id IN (17027, 17038); # a_gnoll_tactician 14
UPDATE npc_types SET loottable_id = :a_patrolling_gnoll_lt: WHERE id IN (17007, 17004, 17011); # a_patrolling_gnoll 6
UPDATE npc_types SET loottable_id = :fish_lt: WHERE id IN (17025, 17026, 17034, 17052); # a_razorgill 8
UPDATE npc_types SET loottable_id = :a_scrawny_gnoll_lt: WHERE id IN (17000, 17005, 17015); # a_scrawny_gnoll 5
UPDATE npc_types SET loottable_id = :an_elite_gnoll_guard_lt: WHERE id IN (17001, 17020, 17022); # an_elite_gnoll_guard 13
UPDATE npc_types SET loottable_id = :lord_elgnub_lt: WHERE id IN (17029, 17032); # Lord_Elgnub 22
UPDATE npc_types SET loottable_id = :mannan_of_the_sabertooth_lt: WHERE id IN (17050); # Mannan_of_the_Sabertooth 14
UPDATE npc_types SET loottable_id = :master_brewer_lt: WHERE id IN (17049); # Master_Brewer 14
UPDATE npc_types SET loottable_id = :refugee_splitpaw_lt: WHERE id IN (17023); # Refugee_Splitpaw 15
UPDATE npc_types SET loottable_id = :scout_malityn_lt: WHERE id IN (17028); # Scout_Malityn 10
UPDATE npc_types SET loottable_id = :socho_darkpaw_lt: WHERE id IN (17048); # Socho_Darkpaw 14
UPDATE npc_types SET loottable_id = :splitpaw_commander_lt: WHERE id IN (17035); # Splitpaw_Commander 15
UPDATE npc_types SET loottable_id = :the_gnoll_high_shaman_lt: WHERE id IN (17042); # the_gnoll_high_shaman 15
UPDATE npc_types SET loottable_id = :tranixx_darkpaw_lt: WHERE id IN (17051); # Tranixx_Darkpaw 17

#--
UPDATE spawn2 SET min_expansion = 5 WHERE spawngroupid = 141; # a_gnoll_courier doesn't spawn until loy
#--
