module main

go 1.22.3

replace ypmodule => ../ypmodule

replace somemodule => ../somemodule

replace go-basics_yandex-practicum/02/03-packages-and-modules/03-modules/somemodule => ../somemodule

require (
	go-basics_yandex-practicum/02/03-packages-and-modules/03-modules/somemodule v0.0.0-00010101000000-000000000000
	ypmodule v0.0.0-00010101000000-000000000000
)
