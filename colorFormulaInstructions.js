// Variables for body R G B values
var bodyRed;
var bodyBlue;
var bodyGreen;

// Variables for body H S B values
var bodyHue;
var bodySaturation;
var bodyBrightness;

// Variables for hair H S B values
var hairHue;
var hairSaturation;
var hairBrightness;

// Min and max perceivedBrightness values (between 0 and 100)
var minPerceivedBrightness = 20;
var maxPerceivedBrightness = 95;

// Min and max perceivedBrightness values (between 0 and 255)
var minPerceivedBrightness255 = minPerceivedBrightness / 100 * 255;
var maxPerceivedBrightness255 = maxPerceivedBrightness / 100 * 255;

// Variable for body and hair hue distance
var bodyAndHairHueDistance = 90;

// Min total saturation (bodySaturation + hairSaturation shouldn't be below this value)
var minTotalSaturation = 60;

// Min total brightness
var minTotalBrightness = 120;

// Min hair brightness value
var minHairBrightness = 30;





////////////////////////////////
// BODY COLOR PICKING PROCESS //
////////////////////////////////

// STEP 1 //
// Pick a random bodyRed value to satisfy 0 <= bodyRed <= 255
bodyRed = Math.floor(Math.random() * 256);

// STEP 1 //
// Pick a random bodyGreen value to satisfy 0 <= bodyRed <= 255
bodyGreen = Math.floor(Math.random() * 256);

// STEP 3 //
// Pick a random bodyBlue value between
Math.max(
    Math.sqrt(
        Math.max(
            (minPerceivedBrightness255 * minPerceivedBrightness255 - 0.241 * bodyRed * bodyRed - 0.691 * bodyGreen * bodyGreen) / 0.068,
            0
        )
    ),
    0
)
// and
Math.min(
    Math.sqrt(
        Math.max(
            (maxPerceivedBrightness255 * maxPerceivedBrightness255 - 0.241 * bodyRed * bodyRed - 0.691 * bodyGreen * bodyGreen) / 0.068,
            0
        )
    ),
    255
)

// STEP 4 //
// Convert these RGB values so that we have hairHue, hairSaturation and hairBrightness as well (to be used on hair color picking process)





////////////////////////////////
// HAIR COLOR PICKING PROCESS //
////////////////////////////////

// STEP 1 //
// Pick a random hair hue between
bodyHue - 180 - bodyAndHairHueDistance
// and
bodyHue - 180 + bodyAndHairHueDistance

if (hairHue < 0) {
    hairHue += 360;
}

// STEP 2 //
// Pick a random hair saturation between
Math.max(minTotalSaturation - bodySaturation, 0)
// and
100

// STEP 3 //
// Pick a random hair brightness between
Math.min(minTotalBrightness - bodyBrightness, 100) // When the perceived brightness of body is low enough, hair brightness can end up being more than 100 here, so we're making sure that hair brightness's minimum value never goes above 100
// and
100
