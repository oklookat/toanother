import { MessageError as MsgErr, MessageInfo as MsgInfo } from "../../wailsjs/go/main/App";

export type Config = {
    YandexMusic: {
        Login: string
    }
    Spotify: {
        ID: string,
        Secret: string
    }
}

/** subscribe to event */
export function EventsOn(eventName: string, callback: (...args: any[]) => void) {
    // @ts-ignore
    if (typeof window.runtime === 'undefined') {
        throw Error("Runtime not available")
    }
    // @ts-ignore
    window.runtime?.EventsOn(eventName, callback)
}

/** unsubscribe from event */
export function EventsOff(eventName: string) {
    // @ts-ignore
    if (typeof window.runtime === 'undefined') {
        throw Error("Runtime not available")
    }
    // @ts-ignore
    window.runtime?.EventsOff(eventName)
}

/** open link in browser */
export function openInBrowser(url: string) {
    // @ts-ignore
    if (typeof window.runtime === 'undefined') {
        throw Error("Runtime not available")
    }
    // @ts-ignore
    window.runtime?.BrowserOpenURL(url)
}

/** show info dialog */
export async function MessageInfo(message: string): Promise<void> {
    await MsgInfo(message)
}

/** show message dialog */
export async function MessageError(message: string): Promise<void> {
    await MsgErr(message)
}