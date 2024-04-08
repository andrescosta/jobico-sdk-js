import { Jobico, Level } from './sdk/sdk.js';

const j = new Jobico();

j.Log(Level.Panic,"test")
j.Log(Level.Panic,"test2")
j.ResultAsJSON(0,j.InputAsJSON())
