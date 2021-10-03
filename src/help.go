package main

import "fmt"

func help(gradient bool) {
	if gradient {fmt.Print("\u001B[38;2;222;31;171mi\u001B[38;2;222;31;167mT\u001B[38;2;222;31;163mC\u001B[38;2;222;31;159mP\u001B[38;2;222;31;155m \u001B[38;2;222;31;151mA\u001B[38;2;222;31;147md\u001B[38;2;222;31;143mv\u001B[38;2;222;31;139ma\u001B[38;2;222;31;135mn\u001B[38;2;222;31;131mc\u001B[38;2;222;31;127me\u001B[38;2;222;31;123md\u001B[38;2;222;31;119m \u001B[38;2;222;31;115mN\u001B[38;2;222;31;111me\u001B[38;2;222;31;107mt\u001B[38;2;222;31;103mw\u001B[38;2;222;31;99mo\u001B[38;2;222;31;95mr\u001B[38;2;222;31;91mk\u001B[38;2;222;31;87m \u001B[38;2;222;31;83mP\u001B[38;2;222;31;79mi\u001B[38;2;222;31;75mn\u001B[38;2;222;31;71mg\u001B[38;2;222;31;67me\u001B[38;2;222;31;63mr\u001B[38;2;222;31;59m \u001B[38;2;222;31;55m(\u001B[38;2;222;31;51mc\u001B[38;2;222;31;47m)\u001B[38;2;222;31;43m \u001B[38;2;222;31;39mY\u001B[38;2;222;31;35ma\u001B[38;2;222;31;31mr\u001B[38;2;222;31;27mi\u001B[38;2;222;31;23my\u001B[38;2;222;31;19ma\u001B[38;2;222;31;15m\n\u001B[38;2;222;31;171m \u001B[38;2;222;31;167m \u001B[38;2;222;31;163m \u001B[38;2;222;31;159m \u001B[38;2;222;31;155m \u001B[38;2;222;31;151m \u001B[38;2;222;31;147m \u001B[38;2;222;31;143m \u001B[38;2;222;31;139m \u001B[38;2;222;31;135m \u001B[38;2;222;31;131m \u001B[38;2;222;31;127m \u001B[38;2;222;31;123m \u001B[38;2;222;31;119m \u001B[38;2;222;31;115m \u001B[38;2;222;31;111m \u001B[38;2;222;31;107mI\u001B[38;2;222;31;103mn\u001B[38;2;222;31;99ms\u001B[38;2;222;31;95mp\u001B[38;2;222;31;91mi\u001B[38;2;222;31;87mr\u001B[38;2;222;31;83me\u001B[38;2;222;31;79md\u001B[38;2;222;31;75m \u001B[38;2;222;31;71mb\u001B[38;2;222;31;67my\u001B[38;2;222;31;63m \u001B[38;2;222;31;59mp\u001B[38;2;222;31;55ma\u001B[38;2;222;31;51mp\u001B[38;2;222;31;47mi\u001B[38;2;222;31;43mn\u001B[38;2;222;31;39mg\u001B[38;2;222;31;35m\n\u001B[38;2;222;31;171m\n\u001B[38;2;222;31;171m \u001B[38;2;222;31;168mS\u001B[38;2;222;31;165my\u001B[38;2;222;31;162mn\u001B[38;2;222;31;159mt\u001B[38;2;222;31;156ma\u001B[38;2;222;31;153mx\u001B[38;2;222;31;150m:\u001B[38;2;222;31;147m \u001B[38;2;222;31;144mi\u001B[38;2;222;31;141mt\u001B[38;2;222;31;138mc\u001B[38;2;222;31;135mp\u001B[38;2;222;31;132m \u001B[38;2;222;31;129md\u001B[38;2;222;31;126me\u001B[38;2;222;31;123ms\u001B[38;2;222;31;120mt\u001B[38;2;222;31;117mi\u001B[38;2;222;31;114mn\u001B[38;2;222;31;111ma\u001B[38;2;222;31;108mt\u001B[38;2;222;31;105mi\u001B[38;2;222;31;102mo\u001B[38;2;222;31;99mn\u001B[38;2;222;31;96m \u001B[38;2;222;31;93m[\u001B[38;2;222;31;90mo\u001B[38;2;222;31;87mp\u001B[38;2;222;31;84mt\u001B[38;2;222;31;81mi\u001B[38;2;222;31;78mo\u001B[38;2;222;31;75mn\u001B[38;2;222;31;72ms\u001B[38;2;222;31;69m]\u001B[38;2;222;31;66m \u001B[38;2;222;31;63m<\u001B[38;2;222;31;60mf\u001B[38;2;222;31;57ml\u001B[38;2;222;31;54ma\u001B[38;2;222;31;51mg\u001B[38;2;222;31;48ms\u001B[38;2;222;31;45m>\u001B[38;2;222;31;42m\n\u001B[38;2;222;31;171m \u001B[38;2;225;29;155mO\u001B[38;2;228;27;139mp\u001B[38;2;231;25;123mt\u001B[38;2;234;23;107mi\u001B[38;2;237;21;91mo\u001B[38;2;240;19;75mn\u001B[38;2;243;17;59ms\u001B[38;2;246;15;43m:\u001B[38;2;249;13;27m\n\u001B[38;2;222;31;171m \u001B[38;2;222;31;168m \u001B[38;2;222;31;165m \u001B[38;2;222;31;162m-\u001B[38;2;222;31;159mp\u001B[38;2;222;31;156m \u001B[38;2;222;31;153m \u001B[38;2;222;31;150m \u001B[38;2;222;31;147m \u001B[38;2;222;31;144m \u001B[38;2;222;31;141m \u001B[38;2;222;31;138m \u001B[38;2;222;31;135m=\u001B[38;2;222;31;132m>\u001B[38;2;222;31;129m \u001B[38;2;222;31;126mP\u001B[38;2;222;31;123mi\u001B[38;2;222;31;120mn\u001B[38;2;222;31;117mg\u001B[38;2;222;31;114ms\u001B[38;2;222;31;111m \u001B[38;2;222;31;108mo\u001B[38;2;222;31;105mn\u001B[38;2;222;31;102m \u001B[38;2;222;31;99mc\u001B[38;2;222;31;96me\u001B[38;2;222;31;93mr\u001B[38;2;222;31;90mt\u001B[38;2;222;31;87ma\u001B[38;2;222;31;84mi\u001B[38;2;222;31;81mn\u001B[38;2;222;31;78m \u001B[38;2;222;31;75mT\u001B[38;2;222;31;72mC\u001B[38;2;222;31;69mP\u001B[38;2;222;31;66m \u001B[38;2;222;31;63mP\u001B[38;2;222;31;60mo\u001B[38;2;222;31;57mr\u001B[38;2;222;31;54mt\u001B[38;2;222;31;51m \u001B[38;2;222;31;48m(\u001B[38;2;222;31;45m-\u001B[38;2;222;31;42mp\u001B[38;2;222;31;39m \u001B[38;2;222;31;36m5\u001B[38;2;222;31;33m3\u001B[38;2;222;31;30m)\u001B[38;2;222;31;27m\n\u001B[38;2;222;31;171m \u001B[38;2;222;31;169m \u001B[38;2;222;31;167m \u001B[38;2;222;31;165m-\u001B[38;2;222;31;163mt\u001B[38;2;222;31;161mi\u001B[38;2;222;31;159mm\u001B[38;2;222;31;157me\u001B[38;2;222;31;155mo\u001B[38;2;222;31;153mu\u001B[38;2;222;31;151mt\u001B[38;2;222;31;149m \u001B[38;2;222;31;147m=\u001B[38;2;222;31;145m>\u001B[38;2;222;31;143m \u001B[38;2;222;31;141mS\u001B[38;2;222;31;139me\u001B[38;2;222;31;137mt\u001B[38;2;222;31;135ms\u001B[38;2;222;31;133m \u001B[38;2;222;31;131ms\u001B[38;2;222;31;129mo\u001B[38;2;222;31;127mc\u001B[38;2;222;31;125mk\u001B[38;2;222;31;123me\u001B[38;2;222;31;121mt\u001B[38;2;222;31;119m \u001B[38;2;222;31;117mt\u001B[38;2;222;31;115mi\u001B[38;2;222;31;113mm\u001B[38;2;222;31;111me\u001B[38;2;222;31;109mo\u001B[38;2;222;31;107mu\u001B[38;2;222;31;105mt\u001B[38;2;222;31;103m \u001B[38;2;222;31;101mt\u001B[38;2;222;31;99mc\u001B[38;2;222;31;97mp\u001B[38;2;222;31;95m/\u001B[38;2;222;31;93mi\u001B[38;2;222;31;91mc\u001B[38;2;222;31;89mm\u001B[38;2;222;31;87mp\u001B[38;2;222;31;85m \u001B[38;2;222;31;83m(\u001B[38;2;222;31;81m-\u001B[38;2;222;31;79mt\u001B[38;2;222;31;77mi\u001B[38;2;222;31;75mm\u001B[38;2;222;31;73me\u001B[38;2;222;31;71mo\u001B[38;2;222;31;69mu\u001B[38;2;222;31;67mt\u001B[38;2;222;31;65m \u001B[38;2;222;31;63m1\u001B[38;2;222;31;61m0\u001B[38;2;222;31;59m0\u001B[38;2;222;31;57m0\u001B[38;2;222;31;55m)\u001B[38;2;222;31;53m\n\u001B[38;2;222;31;171m \u001B[38;2;222;31;170m \u001B[38;2;222;31;169m \u001B[38;2;222;31;168m-\u001B[38;2;222;31;167mp\u001B[38;2;222;31;166mt\u001B[38;2;222;31;165mi\u001B[38;2;222;31;164mm\u001B[38;2;222;31;163me\u001B[38;2;222;31;162m \u001B[38;2;222;31;161m \u001B[38;2;222;31;160m \u001B[38;2;222;31;159m=\u001B[38;2;222;31;158m>\u001B[38;2;222;31;157m \u001B[38;2;222;31;156mS\u001B[38;2;222;31;155me\u001B[38;2;222;31;154mt\u001B[38;2;222;31;153ms\u001B[38;2;222;31;152m \u001B[38;2;222;31;151mt\u001B[38;2;222;31;150mh\u001B[38;2;222;31;149me\u001B[38;2;222;31;148m \u001B[38;2;222;31;147md\u001B[38;2;222;31;146me\u001B[38;2;222;31;145ml\u001B[38;2;222;31;144ma\u001B[38;2;222;31;143my\u001B[38;2;222;31;142m \u001B[38;2;222;31;141mb\u001B[38;2;222;31;140me\u001B[38;2;222;31;139mt\u001B[38;2;222;31;138mw\u001B[38;2;222;31;137me\u001B[38;2;222;31;136me\u001B[38;2;222;31;135mn\u001B[38;2;222;31;134m \u001B[38;2;222;31;133mp\u001B[38;2;222;31;132mi\u001B[38;2;222;31;131mn\u001B[38;2;222;31;130mg\u001B[38;2;222;31;129ms\u001B[38;2;222;31;128m \u001B[38;2;222;31;127ml\u001B[38;2;222;31;126mo\u001B[38;2;222;31;125mw\u001B[38;2;222;31;124me\u001B[38;2;222;31;123mr\u001B[38;2;222;31;122m=\u001B[38;2;222;31;121mf\u001B[38;2;222;31;120ma\u001B[38;2;222;31;119ms\u001B[38;2;222;31;118mt\u001B[38;2;222;31;117me\u001B[38;2;222;31;116mr\u001B[38;2;222;31;115m \u001B[38;2;222;31;114mD\u001B[38;2;222;31;113me\u001B[38;2;222;31;112mf\u001B[38;2;222;31;111ma\u001B[38;2;222;31;110mu\u001B[38;2;222;31;109ml\u001B[38;2;222;31;108mt\u001B[38;2;222;31;107m:\u001B[38;2;222;31;106m \u001B[38;2;222;31;105m1\u001B[38;2;222;31;104m0\u001B[38;2;222;31;103m0\u001B[38;2;222;31;102m0\u001B[38;2;222;31;101m \u001B[38;2;222;31;100m(\u001B[38;2;222;31;99m-\u001B[38;2;222;31;98mp\u001B[38;2;222;31;97mt\u001B[38;2;222;31;96mi\u001B[38;2;222;31;95mm\u001B[38;2;222;31;94me\u001B[38;2;222;31;93m \u001B[38;2;222;31;92m5\u001B[38;2;222;31;91m0\u001B[38;2;222;31;90m0\u001B[38;2;222;31;89m)\u001B[38;2;222;31;88m\n\u001B[38;2;222;31;171m \u001B[38;2;222;31;169m \u001B[38;2;222;31;167m \u001B[38;2;222;31;165m-\u001B[38;2;222;31;163mt\u001B[38;2;222;31;161mh\u001B[38;2;222;31;159mr\u001B[38;2;222;31;157me\u001B[38;2;222;31;155ma\u001B[38;2;222;31;153md\u001B[38;2;222;31;151ms\u001B[38;2;222;31;149m \u001B[38;2;222;31;147m=\u001B[38;2;222;31;145m>\u001B[38;2;222;31;143m \u001B[38;2;222;31;141mM\u001B[38;2;222;31;139ma\u001B[38;2;222;31;137mk\u001B[38;2;222;31;135me\u001B[38;2;222;31;133m \u001B[38;2;222;31;131mm\u001B[38;2;222;31;129mu\u001B[38;2;222;31;127ml\u001B[38;2;222;31;125mt\u001B[38;2;222;31;123mi\u001B[38;2;222;31;121mp\u001B[38;2;222;31;119ml\u001B[38;2;222;31;117me\u001B[38;2;222;31;115m \u001B[38;2;222;31;113mt\u001B[38;2;222;31;111mh\u001B[38;2;222;31;109mr\u001B[38;2;222;31;107me\u001B[38;2;222;31;105ma\u001B[38;2;222;31;103md\u001B[38;2;222;31;101ms\u001B[38;2;222;31;99m \u001B[38;2;222;31;97mf\u001B[38;2;222;31;95mo\u001B[38;2;222;31;93mr\u001B[38;2;222;31;91m \u001B[38;2;222;31;89mp\u001B[38;2;222;31;87mi\u001B[38;2;222;31;85mn\u001B[38;2;222;31;83mg\u001B[38;2;222;31;81mi\u001B[38;2;222;31;79mn\u001B[38;2;222;31;77mg\u001B[38;2;222;31;75m \u001B[38;2;222;31;73mD\u001B[38;2;222;31;71me\u001B[38;2;222;31;69mf\u001B[38;2;222;31;67ma\u001B[38;2;222;31;65mu\u001B[38;2;222;31;63ml\u001B[38;2;222;31;61mt\u001B[38;2;222;31;59m:\u001B[38;2;222;31;57m \u001B[38;2;222;31;55m1\u001B[38;2;222;31;53m \u001B[38;2;222;31;51m(\u001B[38;2;222;31;49m-\u001B[38;2;222;31;47mt\u001B[38;2;222;31;45mh\u001B[38;2;222;31;43mr\u001B[38;2;222;31;41me\u001B[38;2;222;31;39ma\u001B[38;2;222;31;37md\u001B[38;2;222;31;35ms\u001B[38;2;222;31;33m \u001B[38;2;222;31;31m2\u001B[38;2;222;31;29m)\u001B[38;2;222;31;27m\n\u001B[38;2;222;31;171m\n\u001B[38;2;222;31;171m \u001B[38;2;226;28;148mF\u001B[38;2;230;25;125ml\u001B[38;2;234;22;102ma\u001B[38;2;238;19;79mg\u001B[38;2;242;16;56ms\u001B[38;2;246;13;33m\n\u001B[38;2;222;31;171m \u001B[38;2;222;31;169m \u001B[38;2;222;31;167m \u001B[38;2;222;31;165m-\u001B[38;2;222;31;163m-\u001B[38;2;222;31;161md\u001B[38;2;222;31;159me\u001B[38;2;222;31;157m \u001B[38;2;222;31;155m \u001B[38;2;222;31;153m \u001B[38;2;222;31;151m \u001B[38;2;222;31;149m \u001B[38;2;222;31;147m=\u001B[38;2;222;31;145m>\u001B[38;2;222;31;143m \u001B[38;2;222;31;141mD\u001B[38;2;222;31;139mi\u001B[38;2;222;31;137ms\u001B[38;2;222;31;135ma\u001B[38;2;222;31;133mb\u001B[38;2;222;31;131ml\u001B[38;2;222;31;129me\u001B[38;2;222;31;127ms\u001B[38;2;222;31;125m \u001B[38;2;222;31;123mF\u001B[38;2;222;31;121mu\u001B[38;2;222;31;119mn\u001B[38;2;222;31;117mc\u001B[38;2;222;31;115mt\u001B[38;2;222;31;113mi\u001B[38;2;222;31;111mo\u001B[38;2;222;31;109mn\u001B[38;2;222;31;107ms\u001B[38;2;222;31;105m \u001B[38;2;222;31;103ml\u001B[38;2;222;31;101mi\u001B[38;2;222;31;99mk\u001B[38;2;222;31;97me\u001B[38;2;222;31;95m \u001B[38;2;222;31;93mI\u001B[38;2;222;31;91mS\u001B[38;2;222;31;89mP\u001B[38;2;222;31;87m \u001B[38;2;222;31;85ml\u001B[38;2;222;31;83mo\u001B[38;2;222;31;81mo\u001B[38;2;222;31;79mk\u001B[38;2;222;31;77mu\u001B[38;2;222;31;75mp\u001B[38;2;222;31;73m,\u001B[38;2;222;31;71m \u001B[38;2;222;31;69mh\u001B[38;2;222;31;67mo\u001B[38;2;222;31;65ms\u001B[38;2;222;31;63mt\u001B[38;2;222;31;61mn\u001B[38;2;222;31;59ma\u001B[38;2;222;31;57mm\u001B[38;2;222;31;55me\u001B[38;2;222;31;53m \u001B[38;2;222;31;51ml\u001B[38;2;222;31;49mo\u001B[38;2;222;31;47mo\u001B[38;2;222;31;45mk\u001B[38;2;222;31;43mu\u001B[38;2;222;31;41mp\u001B[38;2;222;31;39m\n\u001B[38;2;222;31;171m \u001B[38;2;222;31;167m \u001B[38;2;222;31;163m \u001B[38;2;222;31;159m-\u001B[38;2;222;31;155m-\u001B[38;2;222;31;151mn\u001B[38;2;222;31;147mo\u001B[38;2;222;31;143mi\u001B[38;2;222;31;139mc\u001B[38;2;222;31;135mm\u001B[38;2;222;31;131mp\u001B[38;2;222;31;127m \u001B[38;2;222;31;123m=\u001B[38;2;222;31;119m>\u001B[38;2;222;31;115m \u001B[38;2;222;31;111mD\u001B[38;2;222;31;107mi\u001B[38;2;222;31;103ms\u001B[38;2;222;31;99ma\u001B[38;2;222;31;95mb\u001B[38;2;222;31;91ml\u001B[38;2;222;31;87me\u001B[38;2;222;31;83ms\u001B[38;2;222;31;79m \u001B[38;2;222;31;75mI\u001B[38;2;222;31;71mC\u001B[38;2;222;31;67mM\u001B[38;2;222;31;63mP\u001B[38;2;222;31;59m \u001B[38;2;222;31;55mp\u001B[38;2;222;31;51mi\u001B[38;2;222;31;47mn\u001B[38;2;222;31;43mg\u001B[38;2;222;31;39m\n\u001B[38;2;222;31;171m \u001B[38;2;222;31;167m \u001B[38;2;222;31;163m \u001B[38;2;222;31;159m-\u001B[38;2;222;31;155m-\u001B[38;2;222;31;151mn\u001B[38;2;222;31;147mo\u001B[38;2;222;31;143mt\u001B[38;2;222;31;139mc\u001B[38;2;222;31;135mp\u001B[38;2;222;31;131m \u001B[38;2;222;31;127m \u001B[38;2;222;31;123m=\u001B[38;2;222;31;119m>\u001B[38;2;222;31;115m \u001B[38;2;222;31;111mD\u001B[38;2;222;31;107mi\u001B[38;2;222;31;103ms\u001B[38;2;222;31;99ma\u001B[38;2;222;31;95mb\u001B[38;2;222;31;91ml\u001B[38;2;222;31;87me\u001B[38;2;222;31;83ms\u001B[38;2;222;31;79m \u001B[38;2;222;31;75mT\u001B[38;2;222;31;71mC\u001B[38;2;222;31;67mP\u001B[38;2;222;31;63m \u001B[38;2;222;31;59mp\u001B[38;2;222;31;55mi\u001B[38;2;222;31;51mn\u001B[38;2;222;31;47mg\u001B[38;2;222;31;43m\n\u001B[0;00m")} else {
		fmt.Print(
			"iTCP Advanced Network Pinger (c) Yariya\n",
			"		Inspired by paping\n",
			"\n",
			"Syntax: itcp destination [options] <flags>\n",
			"Options:\n",
			"  -p  		=> Pings on certain TCP Port Required. (-p 53)\n",
			"  -timeout => Sets socket timeout tcp/icmp (-timeout 1000)\n",
			"  -ptime   => Sets the delay between pings lower=faster Default: 1000 (-ptime 500)\n",
			"  -threads => Make multiple threads for pinging Default: 1 (-threads 2)\n",
			"\n",
			"Flags\n",
			"  --de 	=> Disables Functions like ISP lookup, hostname lookup\n",
			"  --noicmp => Disables ICMP ping\n",
			"  --notcp  => Disables TCP ping\n",
		)
	}
}