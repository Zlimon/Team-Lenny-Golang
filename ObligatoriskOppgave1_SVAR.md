**Oppgave 1**

Hvordan gå fra binære eller heksadesimal til desimal:
Bruk formelen verdi = siffer * grunntallposisjon
Grunntallet er alltid det samme som antall tilgjengelige siffer i det tallsystemet,
posisjon er her sifferets posisjon i tallet, posisjonen lengst til høyre er 0, øker mot venstre. 

Eksempel:
10101
1*2^0=1
0*2^1=0
1*2^2=4
0*2^3=0
1*2^4=16
1+0+4+0+16=21

Eksempel:
AF34
4*16^0=4
3*16^1=48
15*16^2=3840
10*16^3=40960
4+48+3840+40960=44852
 
 
Hvordan gå fra desimal til binær:
Del tallet på to, hvis det er rest skriv 1 viss ikke skriv 0
start på toppen helt til svaret blir 0
Hvordan gå fra heksadesimal til binær:
Bruke en tabell, finn det binære 4 sifra tallet for hver karakter 
i det heksadesimale tallet og sett dei sammen 


Eksempel:
50
50/2 = 25, rest 0
25/2 = 12, rest 1
12/2 = 6, rest 0
6/2 = 3, rest 0
3/2 = 1, rest 1
1/2 = 0, rest 1



Eksempel:
ADF
A = 1010
D = 1101
F = 1111
101011011111
 
 
Hvordan gå fra desimal til heksadesimal:
Del tallet på 16, hvis det blir rest skriv det ned, hvis ikke skriv 0.
start på toppen til svaret blir 0
Hvordan gå fra binær til heksadesimal:
del tallet opp i grupper på 4, start bakerst.
fyll ut manglene plasser med ekstra 0-er
 

 
Eksempel:
183
183/16 = 11, rest 7 = B
11/16 = 0, rest 11
B7
Ekempel:
101110
1110 = E
0010 = 2
2E
 
*TABELL* : https://imgur.com/a/uxR0C




**Oppgave 2**

a)

(Se Oppgave2A_BubbleSort.go) 

b)

(Se Oppgave2B_BubbleSort_test.go)

c)

(Se Oppgave2C_QSorting_test.go)



**Oppgave 3**

https://i.imgur.com/o2fkKLy.png
<br>
En avslutningsmelding blir printet når programmet mottar et SIGINT signal:
https://i.imgur.com/n1RN72X.png








**Oppgave 4**:
a)

(Se "Oppgave4A_hexloop2.go" filen)

0x80 - 0x9F blir vist som en firekant istedenfor sitt satte symbol på alle datamaskinene unntatt en macbook hvor 0x80 blir vist som deletegnet. 
Dette er nok på grunn av forskjellig standardvalgt file encoding på windows og macOS.

b)

(Se "Oppgave4B_Oppgave4B.go)












