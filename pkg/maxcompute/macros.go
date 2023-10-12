package maxcompute

import (
	"fmt"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data/sqlutil"
)

var macros = sqlutil.Macros{
	"timeFilter": macroTimeFilter,
	"timeFrom":   macroTimeFrom,
	"timeGroup":  macroTimeGroup,
	"timeTo":     macroTimeTo,
}

func macroTimeFilter(query *sqlutil.Query, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("%w: expected 1 argument, received %d", sqlutil.ErrorBadArgumentCount, len(args))
	}

	var (
		column = args[0]
		from   = query.TimeRange.From.UTC().Format(time.DateTime)
		to     = query.TimeRange.To.UTC().Format(time.DateTime)
	)

	return fmt.Sprintf("%s >= '%s' AND %s <= '%s'", column, from, column, to), nil
}

func macroTimeFrom(query *sqlutil.Query, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("%w: expected 1 argument, received %d", sqlutil.ErrorBadArgumentCount, len(args))
	}

	return fmt.Sprintf("%s >= '%s'", args[0], query.TimeRange.From.UTC().Format(time.DateTime)), nil
}

func macroTimeGroup(_ *sqlutil.Query, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("%w: expected 1 argument, received %d", sqlutil.ErrorBadArgumentCount, len(args))
	}

	column := args[0]

	res := ""
	switch args[1] {
	case "minute":
		res += fmt.Sprintf("datepart('minute', %s) as %s_minute,", column, column)
		fallthrough
	case "hour":
		res += fmt.Sprintf("datepart('hour', %s) as %s_hour,", column, column)
		fallthrough
	case "day":
		res += fmt.Sprintf("datepart('day', %s) as %s_day,", column, column)
		fallthrough
	case "month":
		res += fmt.Sprintf("datepart('month', %s) as %s_month,", column, column)
		fallthrough
	case "year":
		res += fmt.Sprintf("datepart('year', %s) as %s_year", column, column)
	}

	return res, nil
}

func macroTimeTo(query *sqlutil.Query, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("%w: expected 1 argument, received %d", sqlutil.ErrorBadArgumentCount, len(args))
	}

	return fmt.Sprintf("%s <= '%s'", args[0], query.TimeRange.To.UTC().Format(time.DateTime)), nil
}
