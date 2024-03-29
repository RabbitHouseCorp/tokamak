package profilegenerator

import (
	"image"
	"tokamak/src/generator"

	"github.com/fogleman/gg"
)

type ProfileData struct {
	AvatarURL   string   `json:"avatarUrl" form:"avatarUrl"`
	Background  string   `json:"bgId" form:"bgId"`
	Sticker     string   `json:"stickerId" form:"stickerId"` // Deprecated
	Reps        string   `json:"reps" form:"reps"`           // New field
	Married     bool     `json:"married" form:"married"`
	Name        string   `json:"name" form:"name"`
	AboutMe     string   `json:"aboutMe" form:"aboutMe"`
	FavColor    string   `json:"favColor" form:"favColor"`
	Money       string   `json:"money" form:"money"`
	Type        string   `json:"type" form:"type"`
	Badges      []string `json:"badges" form:"badges"`
	PartnerName string   `json:"partnerName" form:"partnerName"`
	AvatarIcon  string   `json:"avatarIcon" form:"avatarIcon"`
}

func RenderDefaultProfile(g generator.Generator, p *ProfileData) image.Image {
	dc := gg.NewContext(600, 400)
	// BASE COLOR
	dc.SetHexColor("212121")
	dc.Clear()

	dc.LoadFontFace("../assets/fonts/Poppins/Poppins-Medium.ttf", 12)

	// BACKGROUND
	img := g.Toolbox.GetAsset("bgs/" + p.Background)
	dc.DrawImage(img, 0, 0)
	rect := gg.NewContext(600, 190)
	rect.SetRGBA(0, 0, 0, 98)
	rect.Clear()
	dc.DrawImage(rect.Image(), 0, 0)

	// R RECT USERNAME
	dc.SetHexColor(p.FavColor)
	dc.DrawRoundedRectangle(180, 160, 170, 35, 5)
	dc.Fill()

	/* STICKER DRAWING. */
	/* WE HAVE TO FIND OUT IF THE USER IS MARRIED. IF THEY ARE,
	WE NEED TO DRAW THE MONEY INDICATOR AND STICKER LOWER ON THE CARD (AFTER MARRIAGE STAT).*/

	// DEF POS (SINGLE)
	stickery := 40
	moneyy := 10

	if p.Married {
		moneyy += 30
		stickery += 30

		// R RECT MARRIAGE
		dc.SetRGBA(1, 1, 1, 170)
		dc.DrawRoundedRectangle(465, 10, 125, 25, 10)
		dc.Fill()

		// RING EMOJI / MARRIED
		img = g.Toolbox.GetAsset("emojis/ring")
		dc.DrawImage(img, 470, 15)

		dc.SetHexColor("000000")
		g.Toolbox.SafeDrawString(dc, p.PartnerName, 490, 26, 105)
	}

	// R RECT YEN
	dc.SetRGBA(1, 1, 1, 170)
	dc.DrawRoundedRectangle(490, float64(moneyy), 100, 25, 10)
	dc.Fill()

	// R RECT STICKER BLUR
	dc.SetRGBA(1, 1, 1, 170)
	dc.DrawRoundedRectangle(490, float64(stickery), 100, 100, 20)
	dc.Fill()
	// R RECT ABOUTME BLUR
	dc.SetRGBA(0, 0, 0, 180)
	dc.DrawRoundedRectangle(350, 225, 225, 150, 20)
	dc.Fill()

	// STICKER DRAWING
	img = g.Toolbox.GetAsset("stickers/" + p.Sticker)
	dc.DrawImage(img, 490, stickery)

	// WHITE CONTOURS
	dc.SetHexColor("ffffff")
	dc.SetLineWidth(6)
	dc.DrawLine(0, 193, 600, 193)
	dc.Stroke()

	/* FAV COLOR DEPENDENT OUTLINING */
	// AVATAR
	dc.SetHexColor(p.FavColor)
	dc.DrawCircle(90, 190, 74)
	dc.Fill()

	// LINE
	dc.SetHexColor(p.FavColor)
	dc.SetLineWidth(5)
	dc.DrawLine(0, 190, 600, 190)
	dc.Stroke()

	// ABOUT ME
	dc.SetHexColor(p.FavColor)
	dc.DrawRoundedRectangle(365, 217, 110, 20, 10)
	dc.Fill()

	// ABOUT ME (EMOJI WHITE CIRCLE)
	dc.SetHexColor("ffffff")
	dc.DrawCircle(381, 227, 11)
	dc.Fill()

	/* EMOJIS (15x15) */
	// YEN/MONEY
	img = g.Toolbox.GetAsset("emojis/money")
	dc.DrawImage(img, 497, moneyy+4)

	/* AVATAR DRAWING */
	avatarSize := 150
	x := 90
	y := 190

	// WHITE OUTLINE @ AVATAR
	dc.SetHexColor("ffffff")
	dc.DrawCircle(float64(x), float64(y), 70)
	dc.Fill()

	avatar := g.Toolbox.ReadImageFromURL(p.AvatarURL, avatarSize, avatarSize)
	dc.DrawCircle(float64(x), float64(y), 66)
	dc.Clip()
	dc.DrawImageAnchored(avatar, x, y, 0.5, 0.5)
	dc.ResetClip()

	// YENS
	dc.SetHexColor("000000")
	dc.DrawString(p.Money, 519, float64(moneyy)+16.506)

	// "ABOUT" ME" @ about ME
	img = g.Toolbox.GetAsset("emojis/woman_laptop")
	dc.DrawImage(img, 373, 219)

	dc.SetHexColor(g.Toolbox.GetCompatibleFontColor(p.FavColor))
	dc.DrawString("About me", 398, 231)

	dc.SetHexColor("ffffff")
	g.Toolbox.DrawTextWrapped(dc, p.AboutMe, 360, 255, 208, 408, 13)

	/* BADGES */
	bx := 25.0
	by := 275.0
	badgesizex := 35.0
	badgesizey := 30.0
	badgespacing := 15.0
	spacebtwedge := 10.0
	recsizex := 175.0

	recsizey := badgesizey*2 + badgespacing
	// R RECT BADGES BLUR
	dc.SetRGBA(0, 0, 0, 180) // 225 150
	dc.DrawRoundedRectangle(bx, by, bx+recsizex, by-recsizex-5, 20)
	dc.Fill()

	cxpos := bx + spacebtwedge
	cypos := by + spacebtwedge

	for _, b := range p.Badges {
		if b != "" {
			dc.DrawImage(g.Toolbox.GetAsset("badges/"+b), int(cxpos), int(cypos))
		}

		cxpos = cxpos + badgesizex + badgespacing
		if cxpos > bx+spacebtwedge+recsizex {
			cypos = cypos + badgesizey + badgespacing
			cxpos = bx + spacebtwedge
			if cypos > by+spacebtwedge+recsizey {
				break // if we run out of space, break the loop
			}
		}
	}

	/* TEXT RENDERING SECTION */
	dc.LoadFontFace("../assets/fonts/Poppins/Poppins-Bold.ttf", 23)

	// USERNAME
	// you're the bird on the brim, hypnotized by the whirl
	dc.SetHexColor(g.Toolbox.GetCompatibleFontColor(p.FavColor))
	g.Toolbox.SafeDrawString(dc, p.Name, 187, 183, 160)

	return dc.Image()
}
