package printing3d

import (
	"math/rand"
	"time"
)

// Tip represents a 3D printing tip
type Tip struct {
	Text     string
	Category string
}

// All 3D printing tips
var tips = []Tip{
	// Bed Adhesion & First Layer
	{Text: "Your first layer is everything. Too high and it won't stick. Too low and you'll scrape the nozzle. Find that sweet spot like Goldilocks.", Category: "First Layer"},
	{Text: "Clean your print bed before every print. Fingerprints contain oils that prevent adhesion. Your prints don't need your personal touch.", Category: "Bed Adhesion"},
	{Text: "Blue painter's tape works great for PLA. It's cheap, easy to replace, and your prints pop right off. Low-tech solutions often win.", Category: "Bed Adhesion"},
	{Text: "Glass beds are flat and easy to clean, but watch that first layer. Add glue stick for PLA or hairspray for PETG if prints won't stick.", Category: "Bed Adhesion"},
	{Text: "Level your bed religiously. An unlevel bed causes more failed prints than any other issue. Take the time. Do it right.", Category: "First Layer"},
	{Text: "Use a brim for small prints or prints with tiny contact points. That extra material around the base prevents warping and gives stability.", Category: "Bed Adhesion"},
	{Text: "Rafts waste filament but save difficult prints. Sometimes throwing material at the problem is the right answer.", Category: "Bed Adhesion"},

	// Temperature & Filament
	{Text: "Print a temperature tower when trying new filament. Different brands print at different temps, even if they're both labeled PLA.", Category: "Temperature"},
	{Text: "PETG is stronger than PLA but strings like a spider on caffeine. Dial in retraction settings or accept your stringy fate.", Category: "Filament"},
	{Text: "Store filament in sealed bags with desiccant. Moisture absorbed from air causes pops, bubbles, and poor layer adhesion. Keep it dry.", Category: "Filament Storage"},
	{Text: "TPU is flexible filament. Print it slow (20-30mm/s) with no retraction. It's like pushing rope through your extruder otherwise.", Category: "Filament"},
	{Text: "ABS requires an enclosure and smells terrible. PLA is easier. PETG is stronger. Choose wisely based on your needs and ventilation.", Category: "Filament"},
	{Text: "Nylon absorbs moisture like a sponge. Dry it thoroughly before printing or you'll get bubbles and weak layers. A food dehydrator works great.", Category: "Filament Storage"},
	{Text: "White and light-colored filaments hide layer lines better than dark colors. Black shows every imperfection like it's under a microscope.", Category: "Filament"},

	// Print Settings
	{Text: "Slow down your print speed if you're getting quality issues. Fast prints look terrible. Slow prints look good. Pick your priority.", Category: "Speed"},
	{Text: "Increase wall count for stronger parts. Two walls is minimum. Three or four walls make functional parts that won't break immediately.", Category: "Strength"},
	{Text: "Infill over 20% is usually wasteful. Use more walls instead. They're visible on the outside and add more strength than hidden infill.", Category: "Infill"},
	{Text: "Layer height affects print time more than speed. 0.2mm layers print twice as fast as 0.1mm layers. Do the math.", Category: "Layer Height"},
	{Text: "Use supports when overhangs exceed 45 degrees. Physics doesn't care about your desire for support-free prints.", Category: "Supports"},
	{Text: "Tree supports use less material and are easier to remove than normal supports. Your future self will thank you.", Category: "Supports"},
	{Text: "Enable 'support interface' layers. They make supports easier to remove and leave better surface finish on the part.", Category: "Supports"},

	// Troubleshooting
	{Text: "Stringing? Lower temperature, increase retraction distance, or speed up travel moves. Usually it's too hot.", Category: "Troubleshooting"},
	{Text: "Layer shifting mid-print means loose belts, overheating steppers, or you bumped the printer. Tighten belts until they twang like a guitar string.", Category: "Troubleshooting"},
	{Text: "Warping corners? Increase bed temperature, add a brim, or eliminate drafts. ABS and PETG warp more than PLA.", Category: "Troubleshooting"},
	{Text: "Blobs and zits appear where the layer starts. Enable 'random seam position' to scatter them or place seam in a corner to hide it.", Category: "Troubleshooting"},
	{Text: "Under-extrusion shows gaps in top layers. Increase flow rate by 5% or check for partial nozzle clogs.", Category: "Troubleshooting"},
	{Text: "Elephant's foot (bulging first layer) means your nozzle is too close to the bed or bed temp is too high. Back it off a hair.", Category: "Troubleshooting"},
	{Text: "Print not sticking? Clean bed, level it, increase bed temp, slow down first layer, or add adhesive. Usually it's one of those five.", Category: "Troubleshooting"},

	// Maintenance
	{Text: "Clean your nozzle regularly. Burnt filament residue builds up and causes clogs and inconsistent extrusion. Cold pulls work wonders.", Category: "Maintenance"},
	{Text: "Lubricate your Z-axis lead screw every few months. A little grease prevents binding and Z-wobble. Don't overdo it.", Category: "Maintenance"},
	{Text: "Check your bowden tube for wear. A gap between tube and nozzle causes jams. Replace it before it ruins a 20-hour print.", Category: "Maintenance"},
	{Text: "Tighten your belts periodically. They stretch over time. Loose belts cause ringing and layer shifting. Tight belts make happy prints.", Category: "Maintenance"},
	{Text: "Keep spare nozzles on hand. They wear out, clog, and cost $1. There's no excuse for not having backups.", Category: "Maintenance"},

	// Design & Slicing
	{Text: "Design parts with 3D printing in mind. Overhangs, thin walls, and tiny details that work in CAD fail in real life. Think like a printer.", Category: "Design"},
	{Text: "Orient parts to minimize supports and maximize strength. Layer lines are weak points. Don't put stress perpendicular to layers.", Category: "Design"},
	{Text: "Add chamfers to bottom edges instead of fillets. They print without supports and look almost as good. Work smarter, not harder.", Category: "Design"},
	{Text: "Use the 'ironing' feature for smooth top surfaces. It takes extra time but makes flat tops look injection molded.", Category: "Slicing"},
	{Text: "Preview your sliced model layer by layer. Catch problems before you waste filament and time. The slicer shows you exactly what will print.", Category: "Slicing"},
	{Text: "Save your slicer profiles when you dial in good settings. You'll want them again and won't remember what you changed.", Category: "Slicing"},

	// Philosophy & Reality
	{Text: "Every printer is different. Settings that work for someone else are a starting point, not gospel. Tune your machine.", Category: "General"},
	{Text: "3D printing is 10% printing and 90% troubleshooting. Embrace the tinkering or buy injection-molded parts.", Category: "General"},
	{Text: "You'll waste filament on failed prints. It's not a question of if, but when and how much. Budget for failures.", Category: "General"},
	{Text: "Patience is required. A 'quick print' takes hours. If you need it now, 3D printing is the wrong tool.", Category: "General"},
	{Text: "Upgrade your printer only after you've mastered it stock. Upgrades won't fix poor technique or lazy bed leveling.", Category: "General"},
	{Text: "Print practical things or fun things. Printing calibration cubes forever means you own an expensive cube factory.", Category: "General"},
	{Text: "Standard nozzle is 0.4mm. Bigger (0.6mm) prints faster but loses detail. Smaller (0.2mm) adds detail but takes forever. Choose based on your part.", Category: "General"},
}

// GetRandomTip returns a random 3D printing tip
func GetRandomTip() Tip {
	rand.Seed(time.Now().UnixNano())
	return tips[rand.Intn(len(tips))]
}

// GetRandomTipWithSeed returns a random 3D printing tip using a specific seed
func GetRandomTipWithSeed(seed int) Tip {
	rand.Seed(int64(seed))
	return tips[rand.Intn(len(tips))]
}
