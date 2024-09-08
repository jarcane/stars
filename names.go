package main

import (
	"fmt"
	"math/rand"
)

var starNames []string = []string{
	"Achernar", "Aldebaran", "Aldrin", "Alpha-Beta", "Alpha", "Andromedae", "Aquarii", "Archimedes", "Arietis", "Aristotle",
	"Armstrong", "Asimov", "Aurigae", "Beta", "Betelgeuse", "Blue-Cauliflower", "Bracewell", "Cancri", "Canis Majoris",
	"Capricorni", "Cassiopeiae", "Centauri", "Ceti", "Chi", "Clarke", "Copernicus", "COROT", "Cygnus", "Daedalus",
	"Dawn", "Delta", "Deneb", "Draconis", "Dyson", "Electra", "Enzmann", "Epsilon", "Eta", "Feynman", "Fomalhaut",
	"Galileo", "Gamma", "Geminorum", "Giedi", "Graarluh", "Grissom", "Hades", "HAT-P", "Hawking", "Heka", "Hephaestus",
	"Herculis", "Horizons", "Hubble", "Hydrae", "Icarus", "Illthrax", "Iota", "Kaku", "Kappa", "Kepler", "Korathraz", "Kuma",
	"Lambda", "Leonis", "Lockhart", "Longshot", "Lucida Anseris", "Lyrae", "Maia", "Medusa", "Mu", "Nembus", "Newton",
	"Nieuw-Vennep", "Nu", "Okul", "Omega", "Omicron", "Ophiuchi", "Oppenheimer", "Orionis", "P-Orion", "Pegasi", "Phi", "Pi",
	"Pink-Green-Teal", "Pioneer", "Plait", "Plato", "Polaris", "Pollux", "Poseidon", "Procyon", "Psi", "Red-Mauve", "Rho",
	"Rigel", "Rorolaah", "Sagan", "Sagittarii", "Sanger", "Scorpii", "Sigma", "Sirius", "Spica", "Starwisp", "Tau",
	"Tauri", "Theta", "Trianguli", "Tycho", "Tyson", "Ulysses", "Upsilon", "Ursae Majoris", "Vega", "Virginis",
	"Voyager", "WASP", "Webb", "Xi", "Zeta", "Zeus",
}

func randomStarName(num int) string {
	return fmt.Sprintf("%d-%s", num, starNames[rand.Intn(len(starNames))])
}
