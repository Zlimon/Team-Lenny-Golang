# IS-105 - Obligatorisk oppgave 1

Gruppe: **Team Lenny**
Deltagere:
*   **Simon Konglevoll Lønnestad**
*   **Petter Sølvsberg Johannessen**
*   **Magnus Husveg**
*   **Kristian Jul-Larsen**
*   **Marius Evensen**

* * *

## Oppgave 1

**Hvordan gå fra binære eller heksadesimal til desimal:**
Bruk formelen verdi = siffer * grunntallposisjon.
Grunntallet er alltid det samme som antall tilgjengelige siffer i det tallsystemet, posisjon er her sifferets posisjon i tallet, posisjonen lengst til høyre er 0, øker mot venstre.

*Eksempel:*
```
10101
1*2^0=1
0*2^1=0
1*2^2=4
0*2^3=0
1*2^4=16
1+0+4+0+16=21
```

*Eksempel:*
```
AF34
4*16^0=4
3*16^1=48
15*16^2=3840
10*16^3=40960
4+48+3840+40960=44852
```

**Hvordan gå fra desimal til binær:**
Del tallet på to. Hvis det er rest skriv 1, hvis ikke skriv 0.
Start på toppen helt til svaret blir 0.

*Eksempel:*
```
50
50/2 = 25, rest 0
25/2 = 12, rest 1
12/2 = 6, rest 0
6/2 = 3, rest 0
3/2 = 1, rest 1
1/2 = 0, rest 1
```

**Hvordan gå fra heksadesimal til binær:**
Bruk en tabell, finn det binære fire sifret tallet for hver karakter i det heksadesimale tallet, og sett dem sammen.

*Eksempel:*
```
ADF
A = 1010
D = 1101
F = 1111
101011011111
```

![](https://i.imgur.com/PeVJWZk.png)

**Hvordan gå fra desimal til heksadesimal:**
Del tallet på 16. Hvis det blir rest skriv det ned, hvis ikke skriv 0.
Start på toppen til svaret blir 0.

*Eksempel:*
```
183
183/16 = 11, rest 7 = B
11/16 = 0, rest 11
B7
```

**Hvordan gå fra binær til heksadesimal:**
Del tallet opp i grupper på fire, start bakerst.
Fyll ut manglende plasser med ekstra 0'er

*Eksempel:*
```
101110
1110 = E
0010 = 2
2E
```

## Oppgave 2
[Lenke til oppgave 2](/src/oppgave2)

a) [bubblesort.go](src/oppgave2/bubblesort.go)

b) [bubblesort_test.go](src/oppgave2/bubblesort_test.go)

c) [sorting_test.go](src/oppgave2/sorting_test.go)

## Oppgave 3
[Lenke til oppgave 3](/src/oppgave3)

[loop.go](src/oppgave3/loop.go)

## Oppgave 4
[Lenke til oppgave 4](/src/oppgave4)

a) [hexloop.go](src/oppgave4/hexloop.go)

```0x80 - 0x9F``` blir vist som en firekant istedenfor sitt satte symbol på alle datamaskinene unntatt en MacBook hvor ```0x80``` blir vist som deletegnet. 
Dette er nok på grunn av forskjellig standardvalgt filkoding på Windows og MacOS.

b) [hexcode.go](src/oppgave4/hexcode.go)