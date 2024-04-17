import './web_out.css'

declare global {
    interface Document {
        WebApp: any
    }
}

interface Field {
    value(): string

    dataset: DOMStringMap

    setCustomValidity(error: string): void

    setOninput(fn: ((this: GlobalEventHandlers, ev: Event) => any) | null): void
}

function getRadioNodeListField(frm: HTMLFormElement, fieldName: string): Field | undefined {

    const fieldEl = frm.elements.namedItem(fieldName)

    if (fieldEl instanceof RadioNodeList) {

        const fieldsetEl = (fieldEl.item(0) as HTMLInputElement).closest('fieldset')

        if (fieldsetEl instanceof HTMLFieldSetElement) {
            return {
                setOninput: function (fn: ((this: GlobalEventHandlers, ev: Event) => any)) {
                    fieldsetEl.oninput = fn
                },
                dataset: fieldsetEl.dataset,
                setCustomValidity: function (errMsg: string) {
                    fieldEl.forEach(function (item: HTMLInputElement) {
                        item.setCustomValidity(errMsg)
                    })
                },
                value: function () {
                    return JSON.stringify((new FormData(frm)).getAll(fieldName))
                }
            }
        }
    }

    return undefined
}

function getElementField(frm: HTMLFormElement, fieldName: string): Field | undefined {

    const fieldEl = frm.elements.namedItem(fieldName)

    if (fieldEl instanceof HTMLInputElement) {

        return {
            setOninput: function (fn: ((this: GlobalEventHandlers, ev: Event) => any)) {
                fieldEl.oninput = fn
            },
            dataset: fieldEl.dataset,
            setCustomValidity: function (errMsg: string) {
                fieldEl.setCustomValidity(errMsg)

            },
            value: function () {
                return JSON.stringify((new FormData(frm)).getAll(fieldName))
            }
        }
    }

    return undefined
}

function dataInitValueValidation(field: Field, initValue: string): void {
    if (field.dataset.error !== undefined) {
        if (initValue == field.value()) {
            field.setCustomValidity('invalid')
        } else {
            field.setCustomValidity('')
        }
    }
}

document.WebApp = {}
document.WebApp.initInputServerValidation = function (fieldName: string) {
    const frm = document.currentScript.previousElementSibling.closest('form')

    let field: Field | undefined = getRadioNodeListField(frm, fieldName)

    if (!field) {
        field = getElementField(frm, fieldName)
    }

    if (field) {
        const initValue = field.value()

        dataInitValueValidation(field, initValue)

        field.setOninput(function () {
            dataInitValueValidation(field, initValue)
        })
    }
}
