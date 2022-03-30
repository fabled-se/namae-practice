package namer

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
 Function ThemedName takes a theme as input
 and returns a name on tha theme
 */

func ThemedName(theme string) string {

	//Seed the random list, using current microtime
	rand.Seed(time.Now().UnixMicro())

	//Get a list of adjectives and a list of nouns
	adjectives, nouns := ThemeLists(theme)

	//Get the length of the adjectives
	adjectivesLength := len(adjectives)

	//Take a random position in the list between 0 and length of adjectives
	adjectivePosition := rand.Intn(adjectivesLength)

	//TODO: Just 3? Use the full length of nouns!
	nounsPosition := rand.Intn(3)

	//Get the first word: an adjective
	firstWord := adjectives[adjectivePosition]

	//Get the last word: an noun = item or person
	lastWord := nouns[nounsPosition]

	//Return a combined string, example: "happy elf"
	return fmt.Sprintf("%s %s", firstWord, lastWord)



}

//TODO: Add a method to get random names, regardless of theme.

//TODO: Add a method to get random nouns, but with adjectives matching theme

//TODO. Add a method to get random adjectives but with nouns matching themes.

//Based on a theme, get two list, one of adjectives, and one of nouns.
func ThemeLists(theme string) ([]string, []string){

	//Match the theme, example. "fantasy" to two lists
	switch(strings.ToUpper(theme)){
	case "FANTASY":
		return FANTASY_ADJECTIVES,FANTASY_ITEMS
	case "SCIFI":
		return SCIFI_ADJECTIVES,SCIFI_NOUNS
	case "ANIMAL":
		return ADJECTIVES,ANIMALS
	case "PEOPLE":
		return ADJECTIVES, PEOPLE
	case "FANTASY_PEOPLE":
		return FANTASY_ADJECTIVES, FANTASY_PEOPLE
	default: //If you end up here, you made a mistake
		return STUPIDADJECTIVES, STUPIDNOUNS
	}

}



//TODO: Create a separate list of adjectives that apply to animals.

var ADJECTIVES = []string{
	"clumsy","scary","grumpy","evil","good","blue","red",
	"feline","lupine","huge","tiny","medieval","futuristic",
	"japanese","korean","chinese","wary","heinous","purple",
	"inept","apt","vulgar","rude",
	"stinky","vile","poisonous","infectious","amorph","sarcastic",
	"rapid","quick","speedy","wise","macabre","lean","dystopian",
	"utopian","insane","ferrous","translucent","fluent","waxy",
	"aromatic","pungent","organic","miasmic","ephemeral","volcanic",
	"divine","demonic","perplexed","vulpine",
	"outlandish","minimalistic","altruistic","aberrant","deviant","defiant",
	"familiar","quaint","anomalous","preternatural",
}

//TODO: Add an animal
var ANIMALS = []string{
	"cat","fox","wolf","dog","rabbit","elephant","slouch",
	"koala","panda","mantis","snake","cobra","bandicoot",
	"tiger","lion", "liger","elk","moose","cow","reindeer",
	"caribou","giraffe","hyena","stork","eagle","hawk",
	"crow","penguin","bee","caterpillar","beetle","firefly",
	"fly","python","pig","boar","bear","shark","tuna","sailfish",
	"anglerfish","monkfish", "cod","pike", "zander","goose","chicken",
	"hen","quail","sparrow","nightingale","tick","tapeworm","sandworm",
	"jellyfish","octopus","squid","oyster","alpaca","aardvark","dolphin",
	"whale","owl","catfish","wasp","spider","butterfly","alligator","adder",
	"crab","crane","dragonfly","eel","parrot","lynx","salamander","tortoise",
	"turtle","lemur","groundhog","rackoon","bat","badger",
}

//TODO: Add some fantasy names
var FANTASY_PEOPLE = []string{
	"dwarf","elf","warlock","sorceror","witch","sorceress","magician",
	"darkblade","blademaster","shadowdancer","darkblade","soulblade",
	"spellcaster","spellslinger","druid","summoner","hobbit","gnome","orc",
	"mage","archer","bowman","swordsman","spellblade","bard","paladin",
	"cleric","troll","vampire","halfling","halfman","ghoul","trickster","thief","rogue",
	"warrior","barbarian","golem","nosferatu","ascendant","fairy",
}


var FANTASY_ADJECTIVES = []string{
	"lonely","cheerful","happy","magical","haunted","demonic","blessed","grim","deadly","fearsome",
	"enchanted","cursed","unlucky","lucky","fabled",
	"fated","secret","hidden","flying","summoned","blasphemous",
	"unholy","holy","light","heavy", "ancient","sentient","talking",
	"singing","screaming","chanting","invisible","poisoned","sharpened",
	"flawless","unremarkable","remarkable","floating",
}

var FANTASY_ITEMS = []string{
	"sword","axe","charm","crown","amulet","grimoire","armor",
	"leggings","palgrins","hauberc","greeves","spear",
	"halberd","chakram","knife","dagger","dirk","pitchfork","hat","lute","fiddle",
	"belt","boots","sandals","blowpipe","shuriken","familiar","watch",
	"relic","artifact","stick","club","mallet","mace","hammer","saucepan",
	"munition","arrow","feather","bag",
}

var SCIFI_ADJECTIVES = [] string {
	"atomic","stellar","interstellar","transcending","ascended","solar",
	"emitting","expanding","quantum","flying","hypersonic","supersonic","ultrasonic",
	"dimensional","interdimensional","alient","extraterrestial",
	"terrestial","teleporing","lightbending","hyperbolic","ionic","descending","orbiting",
	"lunar","planetary","electric","nuclear","guassian","attacking","focused","advanced",
	"ancient","volatile","photovoltaic",
}

var SCIFI_NOUNS = [] string{
	"starship","sun","planet","teleporter","laser","taser","flight","voyage","journey","probe",
	"station","ray","cannon","railgun","missile","robot","android","swarm","flyer","spaceship",
	"spaceman", "being","visitor","asteroid","meteor","barrier","shield","technician","engine",
	"jet","beam","disruptor","bomb","drive","settlement","habitat","civilization","artifact",
}

var PEOPLE = [] string{
	"doctor","nurse","teacher","lawyer","programmer","coder","hacker","mathematician","chemist","pilot",
	"thief","leader","chief","officer","soldier","smith","carpenter","physician","masseur","salesperson",
	"policeman","policewoman","spy","agent","expert","technician","man","woman","baker","florist","veterinary",
	"breeder","beekeeper","captain","priest","preacher","president","inmate","villain","criminal","boss","builder",
	"tester","cook","chef","barista","waiter","waitress","butcher","farmer","brewer", "cashier","astronaut",
	"researcher","lunatic","inventor","lumberjack","fisherman","diver","player","swimmer","gamer","athlete",
}


//Used if you make a mistake...
var STUPIDADJECTIVES = [] string{
	"STUPID","STOOOOPID","IDIOT","MISTAKEN","LOUSY","LAZY","CRAZY","BAT-SHIT-CRAZY","BAD","BAAAAAAAD","HAHAHAHAHA","STINKY",
}

var STUPIDNOUNS = []string{
	"BAKA","LOSER","IDIOT","IMBECILE","FAKE-PROGRAMMER","TIM","SIGGE","PIG","ASSWIPE","BASTARD","SCUMBAG","SHITBRAINS",
}
