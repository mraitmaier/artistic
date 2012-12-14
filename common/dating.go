
package artistic

type Dating int
const (
    Dating_L Dating = iota
    Dating_S
    Dating_A
    Dating_aq
    Dating_aqn
    Dating_pq
    Dating_pqn
    Dating_Unknown
)

func (d Dating) String() string {
    s := ""
    switch d {
    case Dating_L:
        s = "L"
    case Dating_S:
        s = "S"
    case Dating_A:
        s = "A"
    case Dating_aq:
        s = "a.q."
    case Dating_aqn:
        s = "a.q.n."
    case Dating_pq:
        s = "p.q."
    case Dating_pqn:
        s = "p.q.n."
    default:
        s = "Unknown dating"
    }
    return s
}

func convDating2str(val Dating) string {
    d := ""
    switch val {
    case Dating_L:
        d = "L"
    case Dating_S:
        d = "S"
    case Dating_A:
        d = "A"
    case Dating_aq:
        d = "a.q."
    case Dating_aqn:
        d = "a.q.n."
    case Dating_pq:
        d = "p.q."
    case Dating_pqn:
        d = "p.q.n."
    default:
        d = "Unknown dating"
    }
    return d
}

func convDating2val(s string) Dating {
    val := Dating_Unknown

    switch s {
    case "L":
        val = Dating_L
    case "S":
        val = Dating_S
    case "A":
        val = Dating_A
    case "a.q.":
        val = Dating_aq
    case "a.q.n.":
        val = Dating_aqn
    case "p.q.":
        val = Dating_pq
    case "p.q.n.":
        val = Dating_pqn
    default:
        val = Dating_Unknown
    }
    return val
}


