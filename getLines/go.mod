module example.moh/getLines

go 1.22.2

replace example.moh/handleArgs => ../handleArgs

replace example.moh/handleFlag => ../handleFlag

require example.moh/handleArgs v0.0.0-00010101000000-000000000000

require example.moh/handleFlag v0.0.0-00010101000000-000000000000 // indirect
