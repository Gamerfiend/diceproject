#Dice​ ​Project 

Design​ ​the​ ​game​ ​described​ ​below.​ ​Implement​ ​a​ ​method​ ​playGame()​ ​which​ ​simulates​ ​a​ ​game​ ​then displays​ ​the​ ​winning​ ​player​ ​and​ ​their​ ​score. 
 Game​ ​rules: In​ ​this​ ​game​ ​players​ ​roll​ ​dice​ ​and​ ​try​ ​to​ ​collect​ ​the​ ​lowest​ ​score.​ ​A​ ​4​ ​counts​ ​as​ ​zero,​ ​all​ ​other numbers​ ​count​ ​as​ ​face​ ​value.​ ​A​ ​player’s​ ​score​ ​is​ ​the​ ​total​ ​spots​ ​showing​ ​on​ ​the​ ​dice​ ​when​ ​she finishes​ ​her​ ​turn​ ​(excluding​ ​fours​ ​because​ ​they’re​ ​zero).​ ​The​ ​object​ ​of​ ​the​ ​game​ ​is​ ​to​ ​get​ ​the​ ​lowest score​ ​possible. 

The​ ​game​ ​is​ ​played​ ​as​ ​follows​ ​between​ ​4​ ​players: 
 * Play​ ​starts​ ​with​ ​one​ ​person​ ​randomly​ ​being​ ​chosen​ ​to​ ​start​ ​rolling​ ​and​ ​proceeds​ ​in succession​ ​until​ ​all​ ​players​ ​have​ ​rolled. 
 * The​ ​player​ ​rolls​ ​all​ ​five​ ​dice,​ ​then​ ​must​ ​keep​ ​at​ ​least​ ​one​ ​dice​ ​but​ ​may​ ​keep​ ​more​ ​at​ ​her discretion​ ​(She​ ​can​ ​stop​ ​on​ ​her​ ​first​ ​roll​ ​if​ ​she​ ​so​ ​wishes). 
 * Those​ ​dice​ ​which​ ​are​ ​not​ ​kept​ ​are​ ​rolled​ ​again​ ​and​ ​each​ ​round​ ​she​ ​must​ ​keep​ ​at​ ​least​ ​one more​ ​until​ ​all​ ​the​ ​dice​ ​are​ ​out. 
 * Once​ ​each​ ​player​ ​has​ ​rolled​ ​the​ ​player​ ​who​ ​scored​ ​the​ ​lowest​ ​wins. 
 * Repeat​ ​for​ ​three​ ​more​ ​rounds​ ​in​ ​succession​ ​so​ ​that​ ​the​ ​next​ ​person​ ​starts​ ​rolling​ ​first​ ​(at the​ ​end​ ​each​ ​player​ ​will​ ​have​ ​started). After​ ​all​ ​four​ ​rounds​ ​have​ ​been​ ​completed​ ​the​ ​player​ ​with​ ​the​ ​lowest​ ​combined​ ​score​ ​wins. 
 