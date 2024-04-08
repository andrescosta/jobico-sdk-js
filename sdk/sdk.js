export const Level = Object.freeze({
    Debug: Symbol(0),
    Info:Symbol(1),
    Warn:Symbol(2),
    Error:Symbol(3),
    Fatal:Symbol(4),
    Panic:Symbol(5),
    NoLevel:Symbol(6),
})


export class Jobico{
    Log(lvl,desc){
        printErr(lvl.description+desc)
    }
    Result(res, desc){
        print(this.#forceInt32ToString(res)+desc)
    }
    ResultAsJSON(res, desc){
        let resStr = this.#forceInt32ToString(res)
        let jsonStr = JSON.stringify(desc)
        print(resStr+jsonStr)
    }
    #forceInt32ToString(num) {
        num = num | 0;
        return String(num).padStart(11, ' ');
    }
    Input() {
        return readline().replace(/(\r\n|\n|\r)/gm, "");
    }
    InputAsJSON(){
        let s = this.Input()
        return JSON.parse(s);
    }
}

