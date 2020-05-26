import { blake2bFinal, blake2bInit, blake2bUpdate } from 'blakejs'

import { nacl } from "~/plugins/nacl.js"

function getAddressFromPublic(accountPublicKeyBytes, prefix = "nano") {
    const accountHex = uint8ToHex(accountPublicKeyBytes)
    const keyBytes = uint4ToUint8(hexToUint4(accountHex)) // For some reason here we go from u, to hex, to 4, to 8??
    const context = blake2bInit(5)
    blake2bUpdate(context, keyBytes)
    const checksum = uint5ToString(
        uint4ToUint5(uint8ToUint4(blake2bFinal(context).reverse()))
    )
    const account = uint5ToString(uint4ToUint5(hexToUint4(`0${accountHex}`)))

    return `${prefix}_${account}${checksum}`
}

/** Byte Helpers */

function hexToUint4(hexValue) {
    const uint4 = new Uint8Array(hexValue.length)
    for (let i = 0; i < hexValue.length; i++)
        uint4[i] = parseInt(hexValue.substr(i, 1), 16)

    return uint4
}

function uint8ToUint4(uintValue) {
    const uint4 = new Uint8Array(uintValue.length * 2)
    for (let i = 0; i < uintValue.length; i++) {
        uint4[i * 2] = (uintValue[i] / 16) | 0
        uint4[i * 2 + 1] = uintValue[i] % 16
    }

    return uint4
}

function uint8ToHex(uintValue) {
    let hex = ""
    let aux
    for (let i = 0; i < uintValue.length; i++) {
        aux = uintValue[i].toString(16).toUpperCase()
        if (aux.length == 1) aux = "0" + aux
        hex += aux
        aux = ""
    }

    return hex
}

function uint4ToUint8(uintValue) {
    const length = uintValue.length / 2
    const uint8 = new Uint8Array(length)
    for (let i = 0; i < length; i++)
        uint8[i] = uintValue[i * 2] * 16 + uintValue[i * 2 + 1]

    return uint8
}

function uint4ToUint5(uintValue) {
    var length = (uintValue.length / 5) * 4
    var uint5 = new Uint8Array(length)
    for (let i = 1; i <= length; i++) {
        let n = i - 1
        let m = i % 4
        let z = n + (i - m) / 4
        let right = uintValue[z] << m
        let left
        if ((length - i) % 4 == 0) left = uintValue[z - 1] << 4
        else left = uintValue[z + 1] >> (4 - m)
        uint5[n] = (left + right) % 32
    }
    return uint5
}

function uint5ToString(uint5) {
    const letter_list = "13456789abcdefghijkmnopqrstuwxyz".split("")
    let string = ""
    for (let i = 0; i < uint5.length; i++) string += letter_list[uint5[i]]

    return string
}

export function genAddress() {
    const seed = nacl.randomBytes(32);
    const indexBuffer = new ArrayBuffer(4)
    const indexView = new DataView(indexBuffer)
    indexView.setUint32(0, 0)
    const indexBytes = new Uint8Array(indexBuffer)

    const context = blake2bInit(32)
    blake2bUpdate(context, seed)
    blake2bUpdate(context, indexBytes)
    const privateKey = blake2bFinal(context)

    const publicKey = nacl.sign.keyPair.fromSecretKey(privateKey).publicKey
    const address = getAddressFromPublic(publicKey)

    return address
}
