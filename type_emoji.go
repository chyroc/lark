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
	EmojiOk                       EmojiType = "OK"                       // ok https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_ok_v2.png
	EmojiThumbsUp                 EmojiType = "THUMBSUP"                 // 大拇指 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_thumbsup.png
	EmojiThanks                   EmojiType = "THANKS"                   // 谢谢 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_thanks.png
	EmojiMuscle                   EmojiType = "MUSCLE"                   // 肌肉 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fighting.png
	EmojiFingerHeart              EmojiType = "FINGERHEART"              // 比心 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fingerheart.png
	EmojiApplause                 EmojiType = "APPLAUSE"                 // 鼓掌 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_applaud.png
	EmojiFistBump                 EmojiType = "FISTBUMP"                 // 碰拳 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fistbump.png
	EmojiJiaYi                    EmojiType = "JIAYI"                    // +1 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_plusone.png
	EmojiDone                     EmojiType = "DONE"                     // 完成 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_done.png
	EmojiSmile                    EmojiType = "SMILE"                    // 微笑 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smile.png
	EmojiBlush                    EmojiType = "BLUSH"                    // 粘红 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_grin.png
	EmojiLaugh                    EmojiType = "LAUGH"                    // 大笑 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_laugh.png
	EmojiSmirk                    EmojiType = "SMIRK"                    // 傻笑 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smirk.png
	EmojiLol                      EmojiType = "LOL"                      // 笑哭 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lol.png
	EmojiFacepalm                 EmojiType = "FACEPALM"                 // 捂脸 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_facepalm.png
	EmojiLove                     EmojiType = "LOVE"                     // 爱 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_love.png
	EmojiWink                     EmojiType = "WINK"                     // 吐舌 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wink.png
	EmojiProud                    EmojiType = "PROUD"                    // 自豪 https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_proud.png
	EmojiWitty                    EmojiType = "WITTY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_witty.png
	EmojiSmart                    EmojiType = "SMART"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smart.png
	EmojiScowl                    EmojiType = "SCOWL"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_scowl.png
	EmojiThinking                 EmojiType = "THINKING"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_thinking.png
	EmojiSob                      EmojiType = "SOB"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_sob.png
	EmojiCry                      EmojiType = "CRY"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_cry.png
	EmojiError                    EmojiType = "ERROR"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_errr.png
	EmojiNosePick                 EmojiType = "NOSEPICK"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_nosepick.png
	EmojiHaughty                  EmojiType = "HAUGHTY"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_irritated_v2.png
	EmojiSlap                     EmojiType = "SLAP"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_slap.png
	EmojiSpitBlood                EmojiType = "SPITBLOOD"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_spitblood.png
	EmojiToasted                  EmojiType = "TOASTED"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_toasted.png
	EmojiBlackface                EmojiType = "BLACKFACE"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_blackface.png
	EmojiGlance                   EmojiType = "GLANCE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_glance.png
	EmojiDull                     EmojiType = "DULL"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_dull.png
	EmojiRose                     EmojiType = "ROSE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_rose.png
	EmojiHeart                    EmojiType = "HEART"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_heart.png
	EmojiParty                    EmojiType = "PARTY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_party.png
	EmojiInnocentSmile            EmojiType = "INNOCENTSMILE"            // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_innocentsmile.png
	EmojiShy                      EmojiType = "SHY"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_shy.png
	EmojiChuckle                  EmojiType = "CHUCKLE"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_chuckle.png
	EmojiJoyful                   EmojiType = "JOYFUL"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_joyful.png
	EmojiWow                      EmojiType = "WOW"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wow.png
	EmojiTrick                    EmojiType = "TRICK"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_trick.png
	EmojiYeah                     EmojiType = "YEAH"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_yeah.png
	EmojiEnough                   EmojiType = "ENOUGH"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_enough.png
	EmojiTears                    EmojiType = "TEARS"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_tears.png
	EmojiEmbarrassed              EmojiType = "EMBARRASSED"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_embarrassed.png
	EmojiKiss                     EmojiType = "KISS"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_kiss.png
	EmojiSmooch                   EmojiType = "SMOOCH"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smooch.png
	EmojiDrool                    EmojiType = "DROOL"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_drool.png
	EmojiObsessed                 EmojiType = "OBSESSED"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_obsessed.png
	EmojiMoney                    EmojiType = "MONEY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_dollar.png
	EmojiTease                    EmojiType = "TEASE"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_tease.png
	EmojiShowoff                  EmojiType = "SHOWOFF"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_showoff.png
	EmojiComfort                  EmojiType = "COMFORT"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_comfort.png
	EmojiClap                     EmojiType = "CLAP"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_clap.png
	EmojiPraise                   EmojiType = "PRAISE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_praise.png
	EmojiStrive                   EmojiType = "STRIVE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_strive.png
	EmojiXBlush                   EmojiType = "XBLUSH"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_blush.png
	EmojiSilent                   EmojiType = "SILENT"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_silent.png
	EmojiWave                     EmojiType = "WAVE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wave.png
	EmojiEating                   EmojiType = "EATING"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_eating.png
	EmojiWhat                     EmojiType = "WHAT"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_what.png
	EmojiFrown                    EmojiType = "FROWN"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_frown.png
	EmojiDullStare                EmojiType = "DULLSTARE"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_dullstare.png
	EmojiDizzy                    EmojiType = "DIZZY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_dizzy.png
	EmojiLookDown                 EmojiType = "LOOKDOWN"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lookdown.png
	EmojiWail                     EmojiType = "WAIL"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wail.png
	EmojiCrazy                    EmojiType = "CRAZY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_crazy.png
	EmojiWhimper                  EmojiType = "WHIMPER"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_whimper.png
	EmojiHug                      EmojiType = "HUG"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_hug.png
	EmojiBlubber                  EmojiType = "BLUBBER"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_blubber.png
	EmojiWronged                  EmojiType = "WRONGED"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_wronged.png
	EmojiHusky                    EmojiType = "HUSKY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_husky.png
	EmojiShhh                     EmojiType = "SHHH"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_shhh.png
	EmojiSmug                     EmojiType = "SMUG"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_smug.png
	EmojiAngry                    EmojiType = "ANGRY"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_angry.png
	EmojiHammer                   EmojiType = "HAMMER"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_hammer.png
	EmojiShocked                  EmojiType = "SHOCKED"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_shocked.png
	EmojiTerror                   EmojiType = "TERROR"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_terror.png
	EmojiPetrified                EmojiType = "PETRIFIED"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_petrified.png
	EmojiSkull                    EmojiType = "SKULL"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_skull.png
	EmojiSweat                    EmojiType = "SWEAT"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_sweat.png
	EmojiSpeechless               EmojiType = "SPEECHLESS"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_speechless.png
	EmojiSleep                    EmojiType = "SLEEP"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_sleep.png
	EmojiDrowsy                   EmojiType = "DROWSY"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_drowsy.png
	EmojiYawn                     EmojiType = "YAWN"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_yawn.png
	EmojiSick                     EmojiType = "SICK"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_sick.png
	EmojiPuke                     EmojiType = "PUKE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_puke.png
	EmojiBigKiss                  EmojiType = "BIGKISS"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bigkiss.png
	EmojiBetrayed                 EmojiType = "BETRAYED"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_mad.png
	EmojiHeadset                  EmojiType = "HEADSET"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_headset_v3.png
	EmojiDonnotGo                 EmojiType = "DONNOTGO"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_donnotgo.png
	EmojiEatingFood               EmojiType = "EatingFood"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_eatingfood_v2.png
	EmojiTyping                   EmojiType = "Typing"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_typing.png
	EmojiLemon                    EmojiType = "Lemon"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lemon_v2.png
	EmojiGet                      EmojiType = "Get"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_get_v3.png
	EmojiLgtm                     EmojiType = "LGTM"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lgtm_v3.png
	EmojiSalute                   EmojiType = "SALUTE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_salute.png
	EmojiShake                    EmojiType = "SHAKE"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_shake.png
	EmojiHighFive                 EmojiType = "HIGHFIVE"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_highfive.png
	EmojiUpperLeft                EmojiType = "UPPERLEFT"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_upperleft.png
	EmojiThumbsDown               EmojiType = "ThumbsDown"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_thumbsdown.png
	EmojiSlight                   EmojiType = "SLIGHT"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_slight.png
	EmojiTongue                   EmojiType = "TONGUE"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_tongue.png
	EmojiEyesClosed               EmojiType = "EYESCLOSED"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_eyesclosed.png
	EmojiBear                     EmojiType = "BEAR"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bear.png
	EmojiBull                     EmojiType = "BULL"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bull.png
	EmojiCalf                     EmojiType = "CALF"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_calf.png
	EmojiLips                     EmojiType = "LIPS"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_lips.png
	EmojiBeer                     EmojiType = "BEER"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_beer.png
	EmojiCake                     EmojiType = "CAKE"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_cake_v2.png
	EmojiGift                     EmojiType = "GIFT"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_gift_v2.png
	EmojiCucumber                 EmojiType = "CUCUMBER"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_cucumber_v3.png
	EmojiDrumstick                EmojiType = "Drumstick"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_drumstick.png
	EmojiPepper                   EmojiType = "Pepper"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_pepper.png
	EmojiCandiedHaws              EmojiType = "CANDIEDHAWS"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_candiedhaws.png
	EmojiBubbleTea                EmojiType = "BubbleTea"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bubbletea.png
	EmojiCoffee                   EmojiType = "Coffee"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_coffee_v2.png
	EmojiYes                      EmojiType = "Yes"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_yes_v2.png
	EmojiNo                       EmojiType = "No"                       // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_no_v3.png
	EmojiOkr                      EmojiType = "OKR"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_okr.png
	EmojiCheckMark                EmojiType = "CheckMark"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_checkmark.png
	EmojiCrossMark                EmojiType = "CrossMark"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_crossmark.png
	EmojiMinusOne                 EmojiType = "MinusOne"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_minusone_v4.png
	EmojiHundred                  EmojiType = "Hundred"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_hundred.png
	EmojiE2021                    EmojiType = "2021"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_balloon_v2.png
	EmojiAwesome                  EmojiType = "AWESOMEN"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_socool_v2.png
	EmojiPin                      EmojiType = "Pin"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_pin.png
	EmojiAlarm                    EmojiType = "Alarm"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_alarm.png
	EmojiLoudspeaker              EmojiType = "Loudspeaker"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_loudspeaker.png
	EmojiTrophy                   EmojiType = "Trophy"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_trophy.png
	EmojiFire                     EmojiType = "Fire"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fire.png
	EmojiRainbowPuke              EmojiType = "RAINBOWPUKE"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_rainbowpuke_v2.png
	EmojiMusic                    EmojiType = "Music"                    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_music_v2.png
	EmojiFireworks                EmojiType = "FIREWORKS"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_fireworks_v2.png
	EmojiRedPacket                EmojiType = "REDPACKET"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_redpacket_v2.png
	EmojiFortune                  EmojiType = "FORTUNE"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_gold.png
	EmojiLuck                     EmojiType = "LUCK"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_luck.png
	EmojiFirecracker              EmojiType = "FIRECRACKER"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_firecracker.png
	EmojiHeartbroken              EmojiType = "HEARTBROKEN"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_heartbroken.png
	EmojiBomb                     EmojiType = "BOMB"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_bomb.png
	EmojiPoop                     EmojiType = "POOP"                     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_poop.png
	EmojiE18x                     EmojiType = "18X"                      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_eighteen.png
	EmojiCleaver                  EmojiType = "CLEAVER"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_cleaver_v2.png
	EmojiSoccer                   EmojiType = "Soccer"                   // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_soccer_v3.png
	EmojiBasketball               EmojiType = "Basketball"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_basketball_v3.png
	EmojiGeneralDoNotDisturb      EmojiType = "GeneralDoNotDisturb"      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generaldonotdisturb_v2.png
	EmojiStatusPrivateMessage     EmojiType = "Status_PrivateMessage"    // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_status_privatemessage.png
	EmojiGeneralInMeetingBusy     EmojiType = "GeneralInMeetingBusy"     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalInmeetingbusy_v2.png
	EmojiStatusReading            EmojiType = "StatusReading"            // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusreading.png
	EmojiStatusFlashOfInspiration EmojiType = "StatusFlashOfInspiration" // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusflashofInspiration.png
	EmojiGeneralBusinessTrip      EmojiType = "GeneralBusinessTrip"      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalbusinesstrip_v2.png
	EmojiGeneralWorkFromHome      EmojiType = "GeneralWorkFromHome"      // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalworkfromhome.png
	EmojiStatusEnjoyLife          EmojiType = "StatusEnjoyLife"          // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusenjoylife.png
	EmojiGeneralTravellingCar     EmojiType = "GeneralTravellingCar"     // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generaltravellingcar.png
	EmojiStatusBus                EmojiType = "StatusBus"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusbus.png
	EmojiStatusInFlight           EmojiType = "StatusInFlight"           // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_statusinflight.png
	EmojiGeneralSun               EmojiType = "GeneralSun"               // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalsun.png
	EmojiGeneralMoonRest          EmojiType = "GeneralMoonRest"          // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_generalmoonrest_v2.png
	EmojiPursueUltimate           EmojiType = "PursueUltimate"           // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_pursueultimate.png
	EmojiCustomerSuccess          EmojiType = "CustomerSuccess"          // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_customersuccess.png
	EmojiResponsible              EmojiType = "Responsible"              // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_responsible.png
	EmojiReliable                 EmojiType = "Reliable"                 // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_reliable.png
	EmojiAmbitious                EmojiType = "Ambitious"                // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_ambitious.png
	EmojiPatient                  EmojiType = "Patient"                  // https://sf3-ttcdn-tos.pstatp.com/obj/lark-reaction-cn/emoji_patient.png
)
