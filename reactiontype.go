package excel

//go:generate enumer --json --type=ReactionType -trimprefix=Reaction $GOFILE
type ReactionType uint32

const (
	ReactionNone                     ReactionType = 0
	ReactionExplode                  ReactionType = 1
	ReactionStream                   ReactionType = 2
	ReactionBurning                  ReactionType = 3
	ReactionBurned                   ReactionType = 4
	ReactionWet                      ReactionType = 5
	ReactionOvergrow                 ReactionType = 6
	ReactionMelt                     ReactionType = 7
	ReactionFreeze                   ReactionType = 8
	ReactionAntiFire                 ReactionType = 9
	ReactionRock                     ReactionType = 10
	ReactionSlowDown                 ReactionType = 11
	ReactionShock                    ReactionType = 12
	ReactionWind                     ReactionType = 13
	ReactionElectric                 ReactionType = 14
	ReactionFire                     ReactionType = 15
	ReactionSuperconductor           ReactionType = 16
	ReactionSwirlFire                ReactionType = 17
	ReactionSwirlWater               ReactionType = 18
	ReactionSwirlElectric            ReactionType = 19
	ReactionSwirlIce                 ReactionType = 20
	ReactionSwirlFireAccu            ReactionType = 21
	ReactionSwirlWaterAccu           ReactionType = 22
	ReactionSwirlElectricAccu        ReactionType = 23
	ReactionSwirlIceAccu             ReactionType = 24
	ReactionStickRock                ReactionType = 25
	ReactionStickWater               ReactionType = 26
	ReactionCrystallizeFire          ReactionType = 27
	ReactionCrystallizeWater         ReactionType = 28
	ReactionCrystallizeElectric      ReactionType = 29
	ReactionCrystallizeIce           ReactionType = 30
	ReactionFrozenBroken             ReactionType = 31
	ReactionStickGrass               ReactionType = 32
	ReactionOverdose                 ReactionType = 33
	ReactionOverdoseElectric         ReactionType = 34
	ReactionOverdoseGrass            ReactionType = 35
	ReactionOvergrowMushroomFire     ReactionType = 36
	ReactionOvergrowMushroomElectric ReactionType = 37
)
