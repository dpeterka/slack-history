package printing3d

import (
	"github.com/dpeterka/history-slackbot/internal/rotation"
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

	// 2026 refresh
	{Text: "Calibrate your e-steps before touching anything else. If your extruder pushes 95mm when asked for 100, every other setting is a lie.", Category: "Calibration"},
	{Text: "A $20 pair of calipers will improve your prints more than a $200 upgrade. Measure your filament, measure your parts, trust nothing.", Category: "Calibration"},
	{Text: "First layer too squished on one side only? Your gantry is racked. Level the X-axis against the frame with the power off.", Category: "Troubleshooting"},
	{Text: "Hairspray, glue stick, or textured PEI — pick one adhesion method and learn it. Switching daily means you never learn what actually failed.", Category: "Bed Adhesion"},
	{Text: "Dry your filament even if it's 'new.' It sat in a warehouse, on a boat, and in a delivery van. The vacuum seal is a suggestion.", Category: "Filament Storage"},
	{Text: "Wet PETG sounds like bacon frying as it prints. If your printer sounds delicious, your filament needs drying.", Category: "Filament Storage"},
	{Text: "Print your spool holder upgrades in PETG, not PLA. PLA creeps under constant load and your beautiful bracket will slowly sag into modern art.", Category: "Filament"},
	{Text: "Silk PLA looks gorgeous and prints like it hates you. Slow it down, hotten it up, and never use it for functional parts — the layers barely bond.", Category: "Filament"},
	{Text: "Gyroid infill is stronger in all directions, looks cool through translucent walls, and sounds like a UFO landing. There is no downside.", Category: "Infill"},
	{Text: "Set your infill to 100% exactly once, print a chess piece, feel how heavy it is, then never do it again.", Category: "Infill"},
	{Text: "Seams go in corners. A seam on a flat wall is a scar; a seam in a corner is invisible. Tell your slicer where to hide the evidence.", Category: "Slicing"},
	{Text: "Variable layer height is free quality: fine layers on curves, thick layers on straight walls. Two clicks in the slicer, looks twice as good.", Category: "Slicing"},
	{Text: "Your slicer's 'estimated time' is a work of speculative fiction. Add 15% and you'll be pleasantly surprised instead of bitterly disappointed.", Category: "Slicing"},
	{Text: "Fuzzy skin mode hides layer lines completely and makes parts grippy. It's the witness protection program for mediocre print quality.", Category: "Slicing"},
	{Text: "Print small parts in pairs even if you need one. The second copy gives each layer time to cool, and you get a spare for when you drop the first one.", Category: "Speed"},
	{Text: "Speed is free until it isn't. Ringing, ghosting, and skipped steps all arrive at the same party. Find your printer's limit, then back off 20%.", Category: "Speed"},
	{Text: "A clogged nozzle rarely announces itself. Watch your first layer: thin, inconsistent lines mean a partial clog is already taxing your extrusion.", Category: "Troubleshooting"},
	{Text: "Heat creep kills mid-print. If long prints fail at hour three but short ones succeed, your hotend fan is losing the war. Check it, clean it, replace it.", Category: "Troubleshooting"},
	{Text: "Spaghetti detection is cheaper than filament. A $30 camera watching your print saves you from waking up to a plastic bird's nest.", Category: "Troubleshooting"},
	{Text: "PLA parts left in a hot car become abstract sculpture. If it lives outside or in a vehicle, print it in PETG, ASA, or accept the consequences.", Category: "Filament"},
	{Text: "Threaded inserts turn a hobby print into hardware. A soldering iron, a brass insert, and suddenly your part survives being screwed together more than once.", Category: "Design"},
	{Text: "Design in tolerances: holes print smaller than modeled, pegs print bigger. 0.2mm clearance for tight fits, 0.4mm for parts that must slide.", Category: "Design"},
	{Text: "Split tall thin parts and print them lying down. Strength comes from printing along the stress line, not from hoping vertical layers hold.", Category: "Design"},
	{Text: "Keep a print graveyard box. Failed prints are calibration data, test material for paint and glue, and a humbling reminder every time you open the drawer.", Category: "General"},
	{Text: "The best printer upgrade is a notebook. Write down what you changed and why, or you'll fix the same problem four times a year forever.", Category: "General"},
}

// GetRandomTip returns today's 3D printing tip
func GetRandomTip() Tip {
	return GetRandomTipWithSeed(0)
}

// GetRandomTipWithSeed returns a 3D printing tip using a specific seed
func GetRandomTipWithSeed(seed int) Tip {
	if seed == 0 {
		now := time.Now()
		seed = now.Year()*10000 + int(now.Month())*100 + now.Day()
	}
	return tips[rotation.PickIndex(len(tips), seed)]
}
