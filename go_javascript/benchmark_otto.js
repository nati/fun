function ShowBox(name) {
  //console.log("started benchmark : " + name);
}

function AddResult(name, result) {
  console.log(name + "\t" + result);
}

function AddError(name, error) {
  // console.log("error: " + name + ": " + error.message);
}

function AddScore(score) {
  //console.log("score: " + score);
}

BenchmarkSuite.config.doWarmup = undefined;
BenchmarkSuite.config.doDeterministic = undefined;

BenchmarkSuite.RunSuites({
  NotifyStart : ShowBox,
  NotifyError : AddError,
  NotifyResult : AddResult,
  NotifyScore : AddScore
},
[]);
