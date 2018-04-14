package cs



func Walk(start string) {
    Cs := CsParser{Type : "-"}
    ts := make(chan Stats, 1)
    stats := Cs.SmartConcurrent(ts, 8, start)
}
