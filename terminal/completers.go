package terminal

import (
	"github.com/chzyer/readline"
	"raven/database"
	)

var MainCompleter = readline.NewPrefixCompleter(
	readline.PcItem("new" ,
		readline.PcItem("scan"),
	),
	readline.PcItem("use",
		readline.PcItemDynamic(database.GetScanCompleters()),
		),
	readline.PcItem("scans"),
)

var ScanCompleter = readline.NewPrefixCompleter(
	readline.PcItem("options"),
	readline.PcItem("back"),
	readline.PcItem("start"),
	readline.PcItem("set",
		readline.PcItem("scan_name"),
		readline.PcItem("company"),
		readline.PcItem("domain"),
		readline.PcItem("pages_number"),

	),

	readline.PcItem("unset",
		readline.PcItem("scan_name"),
		readline.PcItem("company"),
		readline.PcItem("domain"),
	),
)

var ExportCompleter = readline.NewPrefixCompleter(

	readline.PcItem("export"),
	readline.PcItem("set",
		readline.PcItem("domain"),
		readline.PcItem("output"),
		readline.PcItem("format",
			readline.PcItem("{firstname}.{lastname}@{domain}"),
			readline.PcItem("{lastname}.{firstname}@{domain}"),
			readline.PcItem("{firstname}-{lastname}@{domain}"),
			readline.PcItem("{firstname[0]}{lastname}@{domain}"),
			readline.PcItem("{lastname}{firstname[0]}@{domain}"),
			readline.PcItem("{lastname[0]}{firstname}@{domain}"),
			readline.PcItem("{firstname}{lastname[0]}@{domain}"),
			readline.PcItem("ALL"),
			),
		),
	readline.PcItem("options"),
	readline.PcItem("checkpwned"),
	readline.PcItem("back"),


)