import { EventsOff, EventsOn, openInBrowser } from ".";
import type { base } from "../../wailsjs/go/models";
import { active, args, sign } from "../utils/store";

enum EVENT {
    SPOTIFY_AUTH_URL = "ON_SPOTIFY_AUTH_URL",
	NOT_FOUND     = "ON_NOT_FOUND",
	PROCESSING    = "ON_PROCESSING",
	FINISH        = "ON_FINISH",
}

export class Hooks {

    public static Init() {
        EventsOn(EVENT.PROCESSING, (current: number, total: number) => {
            active.set(true)
            const percents = Math.round((current/total) * 100)
            sign.set(`
            <div>ждём...</div>
            <div>${percents}%</div>
            `)
            args.set([current, total])
        });
        EventsOn(EVENT.FINISH, () => {
            active.set(false)
        });
        EventsOn(EVENT.SPOTIFY_AUTH_URL, (url: string) => {
            openInBrowser(url)
        })
    }

    public static Destroy() {
        active.set(false)
        Object.keys(EVENT).map(key => {
            const evt = EVENT[key]
            EventsOff(evt)
        })
    }

}