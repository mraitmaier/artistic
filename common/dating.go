
package artistic

import (
    "fmt"
)

type DatingValue int
const (
    Dating_Unknown DatingValue = iota
    Dating_L
    Dating_S
    Dating_A
    Dating_aq
    Dating_aqn
    Dating_pq
    Dating_pqn
)

func (d DatingValue) String() string {
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

/**
 *
 */
type Dating struct {
    DatingValue
    Description string
}

func (d *Dating) String() string {
    return d.DatingValue.String()
}

func (d *Dating) Display() string {
    s := fmt.Sprintf("%q\n%s\n", d.DatingValue.String(), d.Description)
    return s
}

func convDating2str(val DatingValue) string {
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

func convDating2val(s string) DatingValue {
    val := Dating_Unknown

    switch s {
    case "L":
        val = Dating_L
    case "S":
        val = Dating_S
    case "A":
        val = Dating_A
    case "a.q.", "aq", "a q":
        val = Dating_aq
//    case "aq":
//        val = Dating_aq
    case "a.q.n.", "aqn", "a q n":
        val = Dating_aqn
//    case "aqn":
//        val = Dating_aqn
    case "p.q.", "pq", "p q":
        val = Dating_pq
//    case "pq":
//        val = Dating_pq
    case "p.q.n.", "pqn", "p q n":
        val = Dating_pqn
//    case "pqn.":
//        val = Dating_pqn
    }
    return val
}


