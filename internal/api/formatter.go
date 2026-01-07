package api

import (
	"fmt"
	"strings"
)

func (p *Pokemon) FormatPokemonInfo() string {
	var builder strings.Builder
	capitalizedName := strings.ToUpper(p.Name[:1]) + p.Name[1:]
	fmt.Fprintln(&builder, "Name:", capitalizedName)
	fmt.Fprintln(&builder, "Height:", p.Height)
	fmt.Fprintln(&builder, "Weight:", p.Weight)

	stats := p.Stats

	if len(stats) > 0 {
		fmt.Fprintln(&builder, "Stats:")
		for _, s := range stats {
			fmt.Fprintf(&builder, " -%v: %v\n", s.Stat.Name, s.BaseStat)
		}
	}

	types := p.Types

	if len(types) > 0 {
		fmt.Fprintln(&builder, "Types:")
		for _, t := range types {
			fmt.Fprintf(&builder, " -%v\n", t.Type.Name)
		}
	}

	return builder.String()
}
