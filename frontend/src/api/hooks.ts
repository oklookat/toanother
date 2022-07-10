import { EventsOff, EventsOn, openInBrowser } from ".";
import type { base } from "../../wailsjs/go/models";
import { active, args, sign } from "../utils/store";

const events = [
    "OnFetchFromAPI",
    "OnFetchFromDatabase",
    "OnAddingToDatabase",
    "OnFinish",
    "SPOTIFY_AUTH_URL",
    "OnImport"
]

export class Hooks {

    public static Init() {
        EventsOn(events[0], (current: number, total: number) => {
            active.set(true)
            sign.set(`
            <div>загрузка с API...</div>
            <div>${current}/${total}</div>
            `)
            args.set([current, total])
        });
        EventsOn(events[1], (current: number, total: number) => {
            active.set(true)
            sign.set(`
            <div>получение из БД...</div> 
            <div>${current}/${total}</div> 
            `)
            args.set([current, total])
        });
        EventsOn(events[2], (current: number, total: number) => {
            active.set(true)
            sign.set(`
            <div>добавление в БД...</div> 
            <div>${current}/${total}</div> 
            `)
            args.set([current, total])
        });
        EventsOn(events[3], () => {
            active.set(false)
        });
        EventsOn(events[4], (url: string) => {
            openInBrowser(url)
        })
        EventsOn(events[5], (current: number, total: number, notFound: any[]) => {
            active.set(true)
            args.set([current, total, notFound])
            sign.set(`
            <div>импорт...</div>
            <div>сейчас: ${current}</div>
            <div>всего: ${total}</div>
            <div>нет: ${notFound.length}</div>
            `)
        })
    }

    public static Destroy() {
        active.set(false)
        for (const event of events) {
            EventsOff(event)
        }
    }

}