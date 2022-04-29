// class MarkdownView {
//   constructor(target, content) {
//     this.textarea = target.appendChild(document.createElement("textarea"))
//     this.textarea.value = content
//   }

//   get content() { return this.textarea.value }
//   focus() { this.textarea.focus() }
//   destroy() { this.textarea.remove() }
// }

import { EditorView } from "prosemirror-view"
import { EditorState } from "prosemirror-state"
import { exampleSetup } from "prosemirror-example-setup"
import { keymap } from "prosemirror-keymap";

import { writeAsMarkdownParser } from "./markdownParser"
import { writeAsMarkdownSerializer } from "./markdownSerializer"
import { writeFreelySchema } from "./schema"
import { getMenu } from "./menu"

let $title = document.querySelector('#title')
let $content = document.querySelector('#content')

class ProseMirrorView {
    constructor(target, content) {
        this.view = new EditorView(target, {
            state: EditorState.create({
                doc: function (content) {
                    // console.log('loading '+window.draftKey)
                    let localDraft = localStorage.getItem(window.draftKey);
                    if (localDraft != null) {
                        content = localDraft
                    }
                    if (content.indexOf("# ") === 0) {
                        let eol = content.indexOf("\n");
                        let title = content.substring("# ".length, eol);
                        content = content.substring(eol + "\n\n".length);
                        $title.value = title;
                    }
                    return writeAsMarkdownParser.parse(content)
                }(content),
                plugins: [
                    keymap({
                        "Mod-Enter": () => {
                            document.getElementById("publish").click();
                            return true;
                        },
                        "Mod-k": ()=> {
                            console.log("TODO-link");
                            return true;
                        }
                    }),
                    ...exampleSetup({ schema: writeFreelySchema, menuContent: getMenu() }),

                ]
            }),
            dispatchTransaction(transaction) {
                // console.log('saving to '+window.draftKey)
                const newContent = writeAsMarkdownSerializer.serialize(transaction.doc)
                $content.value = newContent
                console.log({ newContent })
                localStorage.setItem(window.draftKey, function () {
                    let draft = "";
                    if ($title.value != null && $title.value !== "") {
                        draft = "# " + $title.value + "\n\n"
                    }
                    draft += $content.value
                    return draft
                }());
                let newState = this.state.apply(transaction)
                this.updateState(newState)
            }
        })
    }

    get content() {
        return defaultMarkdownSerializer.serialize(this.view.state.doc)
    }
    focus() { this.view.focus() }
    destroy() { this.view.destroy() }
}

let place = document.querySelector("#editor")
let view = new ProseMirrorView(place, $content.value)

// document.querySelectorAll("input[type=radio]").forEach(button => {
//   button.addEventListener("change", () => {
//     if (!button.checked) return
//     let View = button.value == "markdown" ? MarkdownView : ProseMirrorView
//     if (view instanceof View) return
//     let content = view.content
//     view.destroy()
//     view = new View(place, content)
//     view.focus()
//   })
// })
