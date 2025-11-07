package cal

import (
	"io"
	"time"

	"github.com/emersion/go-ical"
	"github.com/google/uuid"
)

type Color string

var (
	// Apple built in colors
	Fuchsia Color = "#FF2968"
	Orange        = "#FF9500"
	Yellow        = "#FFCC00"
	Green         = "#63DA38"
	Cyan          = "#1BADF8"
	Purple        = "#CC73E1"
	Brown         = "#A2845E"
)

type Calendar struct {
	*ical.Calendar
}

func New(name string) Calendar {
	cal := ical.NewCalendar()

	cal.Props.SetText(ical.PropProductID, "-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN")
	cal.Props.SetText(ical.PropVersion, "2.0")
	cal.Props.SetText("X-WR-CALNAME", name)

	return Calendar{cal}
}

func (cal Calendar) NewEvent(t time.Time, name, desc string) error {
	e := ical.NewEvent()
	e.Props.SetDateTime(ical.PropDateTimeStamp, time.Now())

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	e.Props.SetText(ical.PropUID, id.String())

	e.Props.SetDate(ical.PropDateTimeStart, t)
	e.Props.SetText(ical.PropSummary, name)
	e.Props.SetText(ical.PropDescription, desc)

	cal.Children = append(cal.Children, e.Component)
	return nil
}

func (cal Calendar) Color(c Color) {
	cal.Props.SetText("X-APPLE-CALENDAR-COLOR", string(c))
}

func (cal Calendar) Encode(w io.Writer) error {
	enc := ical.NewEncoder(w)
	return enc.Encode(cal.Calendar)
}
