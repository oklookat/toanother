/** first call = store default body 'no-scroll'
 *
 * second call = set no-scroll if not exists
 * 
 * third call = remove no-scroll if body not have it on first call
 * 
 * @returns body scroll toggle function
 */
export function createBodyScrollToggler(): () => void {
    let active = false
    return () => {
        if (typeof document === 'undefined') {
            console.warn("[createBodyScrollToggler] toggle on client-side, not server")
            return
        }

        const classNames = document.body.className.split(" ")

        if (!active) {
            classNames.push("no-scroll")
            document.body.className = classNames.join(" ").trim()
            active = true
            return
        }

        for (let i = 0; i < classNames.length; i++) {
            const className = classNames[i]
            if (!className) {
                delete classNames[i]
                continue
            }
            if (className === "no-scroll") {
                delete classNames[i]
                break
            }
        }
        document.body.className = classNames.join(" ").trim()
        active = false
    };
}

/** select all text in node */
export function selectText(node: Node) {
    if (!node) {
        return
    }
    const range = document.createRange();
    range.selectNode(node);
    window.getSelection().removeAllRanges();
    window.getSelection().addRange(range);
}
