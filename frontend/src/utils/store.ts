import { writable } from "svelte/store";

export const active = writable(false)
export const sign = writable('idle')
export const args = writable([])
