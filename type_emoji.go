/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package lark

type EmojiType string

const (
	Emoji2022                     EmojiType = "2022"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_twentytwentytwo.png?lazyload=true&width=96&height=96
	EmojiAlarm                    EmojiType = "Alarm"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_alarm.png?lazyload=true&width=96&height=96
	EmojiAmbitious                EmojiType = "Ambitious"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_ambitious.png?lazyload=true&width=96&height=96
	EmojiAngry                    EmojiType = "ANGRY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_angry.png?lazyload=true&width=96&height=96
	EmojiApplause                 EmojiType = "APPLAUSE"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_applaud_v2.png?lazyload=true&width=96&height=96
	EmojiAwesome                  EmojiType = "AWESOMEN"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_socool_v2.png?lazyload=true&width=96&height=96
	EmojiBasketball               EmojiType = "Basketball"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_basketball_v3.png?lazyload=true&width=96&height=96
	EmojiBear                     EmojiType = "BEAR"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bear.png?lazyload=true&width=110&height=96
	EmojiBeer                     EmojiType = "BEER"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_beer.png?lazyload=true&width=96&height=96
	EmojiBetrayed                 EmojiType = "BETRAYED"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_mad.png?lazyload=true&width=97&height=97
	EmojiBigKiss                  EmojiType = "BIGKISS"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bigkiss.png?lazyload=true&width=96&height=96
	EmojiBlackface                EmojiType = "BLACKFACE"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_blackface.png?lazyload=true&width=96&height=96
	EmojiBlubber                  EmojiType = "BLUBBER"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_blubber.png?lazyload=true&width=96&height=96
	EmojiBlush                    EmojiType = "BLUSH"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_grin.png?lazyload=true&width=96&height=96
	EmojiBomb                     EmojiType = "BOMB"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bomb.png?lazyload=true&width=96&height=96
	EmojiBubbleTea                EmojiType = "BubbleTea"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bubbletea.png?lazyload=true&width=96&height=96
	EmojiBull                     EmojiType = "BULL"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bull.png?lazyload=true&width=96&height=96
	EmojiCake                     EmojiType = "CAKE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_cake_v2.png?lazyload=true&width=96&height=96
	EmojiCalf                     EmojiType = "CALF"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_calf.png?lazyload=true&width=96&height=96
	EmojiCandiedHaws              EmojiType = "CANDIEDHAWS"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_candiedhaws.png?lazyload=true&width=96&height=96
	EmojiCheckMark                EmojiType = "CheckMark"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_checkmark.png?lazyload=true&width=96&height=96
	EmojiChuckle                  EmojiType = "CHUCKLE"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_chuckle.png?lazyload=true&width=96&height=96
	EmojiClap                     EmojiType = "CLAP"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_clap.png?lazyload=true&width=102&height=96
	EmojiCleaver                  EmojiType = "CLEAVER"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_cleaver_v2.png?lazyload=true&width=96&height=96
	EmojiCoffee                   EmojiType = "Coffee"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_coffee_v2.png?lazyload=true&width=104&height=96
	EmojiComfort                  EmojiType = "COMFORT"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_comfort.png?lazyload=true&width=96&height=96
	EmojiCrazy                    EmojiType = "CRAZY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_crazy.png?lazyload=true&width=96&height=96
	EmojiCrossMark                EmojiType = "CrossMark"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_crossmark.png?lazyload=true&width=96&height=96
	EmojiCry                      EmojiType = "CRY"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_cry.png?lazyload=true&width=96&height=96
	EmojiCucumber                 EmojiType = "CUCUMBER"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_cucumber_v3.png?lazyload=true&width=96&height=96
	EmojiCustomerSuccess          EmojiType = "CustomerSuccess"          // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_customersuccess.png?lazyload=true&width=96&height=96
	EmojiDizzy                    EmojiType = "DIZZY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_dizzy.png?lazyload=true&width=96&height=96
	EmojiDone                     EmojiType = "DONE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_done.png?lazyload=true&width=96&height=96
	EmojiDonnotGo                 EmojiType = "DONNOTGO"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_donnotgo.png?lazyload=true&width=102&height=96
	EmojiDrool                    EmojiType = "DROOL"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_drool.png?lazyload=true&width=96&height=96
	EmojiDrowsy                   EmojiType = "DROWSY"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_drowsy.png?lazyload=true&width=96&height=96
	EmojiDrumstick                EmojiType = "Drumstick"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_drumstick.png?lazyload=true&width=96&height=96
	EmojiDull                     EmojiType = "DULL"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_dull.png?lazyload=true&width=96&height=96
	EmojiDullStare                EmojiType = "DULLSTARE"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_dullstare.png?lazyload=true&width=96&height=96
	EmojiE18x                     EmojiType = "18X"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_eighteen.png?lazyload=true&width=96&height=96
	EmojiEating                   EmojiType = "EATING"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_eating.png?lazyload=true&width=96&height=96
	EmojiEatingFood               EmojiType = "EatingFood"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_eatingfood_v2.png?lazyload=true&width=96&height=96
	EmojiEmbarrassed              EmojiType = "EMBARRASSED"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_embarrassed.png?lazyload=true&width=96&height=96
	EmojiEnough                   EmojiType = "ENOUGH"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_enough.png?lazyload=true&width=108&height=96
	EmojiError                    EmojiType = "ERROR"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_errr.png?lazyload=true&width=96&height=96
	EmojiEyesClosed               EmojiType = "EYESCLOSED"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_eyesclosed.png?lazyload=true&width=96&height=96
	EmojiFacepalm                 EmojiType = "FACEPALM"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_facepalm.png?lazyload=true&width=96&height=96
	EmojiFingerHeart              EmojiType = "FINGERHEART"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fingerheart_v2.png?lazyload=true&width=96&height=96
	EmojiFire                     EmojiType = "Fire"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fire.png?lazyload=true&width=96&height=96
	EmojiFirecracker              EmojiType = "FIRECRACKER"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_firecracker_v3.png?lazyload=true&width=96&height=96
	EmojiFireworks                EmojiType = "FIREWORKS"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fireworks_v2.png?lazyload=true&width=96&height=96
	EmojiFistBump                 EmojiType = "FISTBUMP"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fistbump_v2.png?lazyload=true&width=96&height=96
	EmojiFortune                  EmojiType = "FORTUNE"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_gold.png?lazyload=true&width=96&height=96
	EmojiFrown                    EmojiType = "FROWN"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_frown.png?lazyload=true&width=96&height=96
	EmojiGeneralBusinessTrip      EmojiType = "GeneralBusinessTrip"      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalbusinesstrip_v2.png?lazyload=true&width=96&height=96
	EmojiGeneralDoNotDisturb      EmojiType = "GeneralDoNotDisturb"      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generaldonotdisturb_v2.png?lazyload=true&width=96&height=96
	EmojiGeneralInMeetingBusy     EmojiType = "GeneralInMeetingBusy"     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalInmeetingbusy_v2.png?lazyload=true&width=96&height=96
	EmojiGeneralMoonRest          EmojiType = "GeneralMoonRest"          // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalmoonrest_v2.png?lazyload=true&width=96&height=96
	EmojiGeneralSun               EmojiType = "GeneralSun"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalsun.png?lazyload=true&width=96&height=96
	EmojiGeneralTravellingCar     EmojiType = "GeneralTravellingCar"     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generaltravellingcar.png?lazyload=true&width=96&height=96
	EmojiGeneralWorkFromHome      EmojiType = "GeneralWorkFromHome"      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalworkfromhome.png?lazyload=true&width=96&height=96
	EmojiGet                      EmojiType = "Get"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_get_v3.png?lazyload=true&width=96&height=96
	EmojiGift                     EmojiType = "GIFT"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_gift_v3.png?lazyload=true&width=96&height=96
	EmojiGlance                   EmojiType = "GLANCE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_glance.png?lazyload=true&width=96&height=96
	EmojiHammer                   EmojiType = "HAMMER"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_hammer.png?lazyload=true&width=96&height=96
	EmojiHaughty                  EmojiType = "HAUGHTY"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_irritated_v2.png?lazyload=true&width=96&height=96
	EmojiHeadset                  EmojiType = "HEADSET"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_headset_v3.png?lazyload=true&width=106&height=96
	EmojiHeart                    EmojiType = "HEART"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_heart.png?lazyload=true&width=96&height=96
	EmojiHeartbroken              EmojiType = "HEARTBROKEN"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_heartbroken.png?lazyload=true&width=96&height=96
	EmojiHighFive                 EmojiType = "HIGHFIVE"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_highfive_v2.png?lazyload=true&width=96&height=96
	EmojiHug                      EmojiType = "HUG"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_hug.png?lazyload=true&width=132&height=96
	EmojiHundred                  EmojiType = "Hundred"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_hundred.png?lazyload=true&width=104&height=96
	EmojiHusky                    EmojiType = "HUSKY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_husky.png?lazyload=true&width=96&height=96
	EmojiInnocentSmile            EmojiType = "INNOCENTSMILE"            // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_innocentsmile.png?lazyload=true&width=96&height=96
	EmojiJiaYi                    EmojiType = "JIAYI"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_plusone.png?lazyload=true&width=96&height=96
	EmojiJoyful                   EmojiType = "JOYFUL"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_joyful.png?lazyload=true&width=96&height=96
	EmojiKiss                     EmojiType = "KISS"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_kiss.png?lazyload=true&width=96&height=96
	EmojiLaugh                    EmojiType = "LAUGH"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_laugh.png?lazyload=true&width=96&height=96
	EmojiLemon                    EmojiType = "Lemon"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lemon_v2.png?lazyload=true&width=96&height=96
	EmojiLgtm                     EmojiType = "LGTM"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lgtm_v3.png?lazyload=true&width=97&height=96
	EmojiLips                     EmojiType = "LIPS"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lips.png?lazyload=true&width=96&height=96
	EmojiLol                      EmojiType = "LOL"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lol.png?lazyload=true&width=108&height=96
	EmojiLookDown                 EmojiType = "LOOKDOWN"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lookdown.png?lazyload=true&width=96&height=96
	EmojiLoudspeaker              EmojiType = "Loudspeaker"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_loudspeaker.png?lazyload=true&width=96&height=96
	EmojiLove                     EmojiType = "LOVE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_love.png?lazyload=true&width=96&height=96
	EmojiLuck                     EmojiType = "LUCK"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_luck.png?lazyload=true&width=96&height=96
	EmojiMeMeMe                   EmojiType = "MeMeMe"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_mememe.png?lazyload=true&width=108&height=96
	EmojiMinusOne                 EmojiType = "MinusOne"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_minusone_v4.png?lazyload=true&width=96&height=96
	EmojiMoney                    EmojiType = "MONEY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_dollar.png?lazyload=true&width=97&height=97
	EmojiMuscle                   EmojiType = "MUSCLE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fighton_v2.png?lazyload=true&width=96&height=96
	EmojiMusic                    EmojiType = "Music"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_music_v2.png?lazyload=true&width=96&height=96
	EmojiNo                       EmojiType = "No"                       // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_no_v3.png?lazyload=true&width=104&height=96
	EmojiNosePick                 EmojiType = "NOSEPICK"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_nosepick.png?lazyload=true&width=96&height=96
	EmojiObsessed                 EmojiType = "OBSESSED"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_obsessed.png?lazyload=true&width=96&height=96
	EmojiOk                       EmojiType = "OK"                       // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_ok_v2.png?lazyload=true&width=104&height=96
	EmojiOkr                      EmojiType = "OKR"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_okr.png?lazyload=true&width=104&height=96
	EmojiOneSecond                EmojiType = "OneSecond"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_onesecond.png?lazyload=true&width=108&height=96
	EmojiOnIt                     EmojiType = "OnIt"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_onit.png?lazyload=true&width=120&height=96
	EmojiParty                    EmojiType = "PARTY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_party.png?lazyload=true&width=96&height=96
	EmojiPatient                  EmojiType = "Patient"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_patient.png?lazyload=true&width=96&height=96
	EmojiPepper                   EmojiType = "Pepper"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_pepper.png?lazyload=true&width=96&height=96
	EmojiPetrified                EmojiType = "PETRIFIED"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_petrified.png?lazyload=true&width=96&height=96
	EmojiPin                      EmojiType = "Pin"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_pin.png?lazyload=true&width=96&height=96
	EmojiPoop                     EmojiType = "POOP"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_poop.png?lazyload=true&width=96&height=96
	EmojiPraise                   EmojiType = "PRAISE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_praise.png?lazyload=true&width=103&height=96
	EmojiProud                    EmojiType = "PROUD"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_proud.png?lazyload=true&width=96&height=96
	EmojiPuke                     EmojiType = "PUKE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_puke.png?lazyload=true&width=96&height=96
	EmojiPursueUltimate           EmojiType = "PursueUltimate"           // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_pursueultimate.png?lazyload=true&width=96&height=96
	EmojiRainbowPuke              EmojiType = "RAINBOWPUKE"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_rainbowpuke_v2.png?lazyload=true&width=96&height=96
	EmojiRedPacket                EmojiType = "REDPACKET"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_redpacket_v2.png?lazyload=true&width=96&height=96
	EmojiReliable                 EmojiType = "Reliable"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_reliable.png?lazyload=true&width=96&height=96
	EmojiResponsible              EmojiType = "Responsible"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_responsible.png?lazyload=true&width=96&height=96
	EmojiRoarForYou               EmojiType = "RoarForYou"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_roarforyou.png?lazyload=true&width=132&height=96
	EmojiRose                     EmojiType = "ROSE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_rose.png?lazyload=true&width=96&height=96
	EmojiSalute                   EmojiType = "SALUTE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_salute_v2.png?lazyload=true&width=96&height=96
	EmojiScowl                    EmojiType = "SCOWL"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_scowl.png?lazyload=true&width=96&height=96
	EmojiShake                    EmojiType = "SHAKE"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_shake_v2.png?lazyload=true&width=96&height=96
	EmojiShhh                     EmojiType = "SHHH"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_shhh.png?lazyload=true&width=96&height=96
	EmojiShocked                  EmojiType = "SHOCKED"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_shocked.png?lazyload=true&width=107&height=96
	EmojiShowoff                  EmojiType = "SHOWOFF"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_showoff.png?lazyload=true&width=96&height=96
	EmojiShy                      EmojiType = "SHY"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_shy.png?lazyload=true&width=96&height=96
	EmojiSick                     EmojiType = "SICK"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_sick.png?lazyload=true&width=96&height=96
	EmojiSigh                     EmojiType = "Sigh"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_sigh.png?lazyload=true&width=96&height=96
	EmojiSilent                   EmojiType = "SILENT"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_silent.png?lazyload=true&width=96&height=96
	EmojiSkull                    EmojiType = "SKULL"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_skull.png?lazyload=true&width=96&height=96
	EmojiSlap                     EmojiType = "SLAP"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_slap.png?lazyload=true&width=96&height=96
	EmojiSleep                    EmojiType = "SLEEP"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_sleep.png?lazyload=true&width=96&height=96
	EmojiSlight                   EmojiType = "SLIGHT"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_slight.png?lazyload=true&width=96&height=96
	EmojiSmart                    EmojiType = "SMART"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smart.png?lazyload=true&width=96&height=96
	EmojiSmile                    EmojiType = "SMILE"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smile.png?lazyload=true&width=96&height=96
	EmojiSmirk                    EmojiType = "SMIRK"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smirk.png?lazyload=true&width=97&height=96
	EmojiSmooch                   EmojiType = "SMOOCH"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smooch.png?lazyload=true&width=100&height=96
	EmojiSmug                     EmojiType = "SMUG"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smug.png?lazyload=true&width=96&height=96
	EmojiSnowman                  EmojiType = "Snowman"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_snowman.png?lazyload=true&width=96&height=96
	EmojiSob                      EmojiType = "SOB"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_sob.png?lazyload=true&width=96&height=96
	EmojiSoccer                   EmojiType = "Soccer"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_soccer_v3.png?lazyload=true&width=96&height=96
	EmojiSpeechless               EmojiType = "SPEECHLESS"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_speechless.png?lazyload=true&width=103&height=96
	EmojiSpitBlood                EmojiType = "SPITBLOOD"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_spitblood.png?lazyload=true&width=102&height=96
	EmojiStatusBus                EmojiType = "StatusBus"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusbus.png?lazyload=true&width=96&height=96
	EmojiStatusEnjoyLife          EmojiType = "StatusEnjoyLife"          // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusenjoylife.png?lazyload=true&width=96&height=96
	EmojiStatusFlashOfInspiration EmojiType = "StatusFlashOfInspiration" // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusflashofInspiration.png?lazyload=true&width=96&height=96
	EmojiStatusInFlight           EmojiType = "StatusInFlight"           // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusinflight.png?lazyload=true&width=96&height=96
	EmojiStatusPrivateMessage     EmojiType = "Status_PrivateMessage"    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_status_privatemessage.png?lazyload=true&width=96&height=96
	EmojiStatusReading            EmojiType = "StatusReading"            // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusreading.png?lazyload=true&width=96&height=96
	EmojiStickyRiceBalls          EmojiType = "StickyRiceBalls"          // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_stickyriceballs.png?lazyload=true&width=106&height=96
	EmojiStrive                   EmojiType = "STRIVE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_strive.png?lazyload=true&width=96&height=96
	EmojiSweat                    EmojiType = "SWEAT"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_sweat.png?lazyload=true&width=100&height=96
	EmojiTears                    EmojiType = "TEARS"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_tears.png?lazyload=true&width=96&height=96
	EmojiTease                    EmojiType = "TEASE"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_tease.png?lazyload=true&width=96&height=96
	EmojiTerror                   EmojiType = "TERROR"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_terror.png?lazyload=true&width=96&height=96
	EmojiThanks                   EmojiType = "THANKS"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_thanks_v2.png?lazyload=true&width=96&height=96
	EmojiThinking                 EmojiType = "THINKING"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_thinking.png?lazyload=true&width=96&height=96
	EmojiThumbsDown               EmojiType = "ThumbsDown"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_thumbsdown_v2.png?lazyload=true&width=96&height=96
	EmojiThumbsUp                 EmojiType = "THUMBSUP"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_thumbsup_v2.png?lazyload=true&width=96&height=96
	EmojiToasted                  EmojiType = "TOASTED"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_toasted.png?lazyload=true&width=96&height=96
	EmojiTongue                   EmojiType = "TONGUE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_tongue.png?lazyload=true&width=96&height=96
	EmojiTrick                    EmojiType = "TRICK"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_trick.png?lazyload=true&width=96&height=96
	EmojiTrophy                   EmojiType = "Trophy"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_trophy.png?lazyload=true&width=96&height=96
	EmojiTyping                   EmojiType = "Typing"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_typing.png?lazyload=true&width=96&height=96
	EmojiUpperLeft                EmojiType = "UPPERLEFT"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_click_v2.png?lazyload=true&width=96&height=96
	EmojiWail                     EmojiType = "WAIL"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wail.png?lazyload=true&width=110&height=96
	EmojiWave                     EmojiType = "WAVE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wave.png?lazyload=true&width=99&height=96
	EmojiWhat                     EmojiType = "WHAT"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_what.png?lazyload=true&width=104&height=96
	EmojiWhimper                  EmojiType = "WHIMPER"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_whimper.png?lazyload=true&width=96&height=96
	EmojiWink                     EmojiType = "WINK"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wink.png?lazyload=true&width=96&height=96
	EmojiWitty                    EmojiType = "WITTY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_witty.png?lazyload=true&width=102&height=96
	EmojiWow                      EmojiType = "WOW"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wow.png?lazyload=true&width=96&height=96
	EmojiWronged                  EmojiType = "WRONGED"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wronged.png?lazyload=true&width=96&height=96
	EmojiXBlush                   EmojiType = "XBLUSH"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_blush.png?lazyload=true&width=96&height=96
	EmojiXmasHat                  EmojiType = "XmasHat"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_xmashat.png?lazyload=true&width=96&height=96
	EmojiXmasTree                 EmojiType = "XmasTree"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_xmastree.png?lazyload=true&width=96&height=96
	EmojiYawn                     EmojiType = "YAWN"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_yawn.png?lazyload=true&width=96&height=96
	EmojiYeah                     EmojiType = "YEAH"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_yeah.png?lazyload=true&width=102&height=96
	EmojiYes                      EmojiType = "Yes"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_yes_v2.png?lazyload=true&width=104&height=96
	EmojiYouAreTheBest            EmojiType = "YouAreTheBest"            // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_youarethebest.png?lazyload=true&width=96&height=96
)
