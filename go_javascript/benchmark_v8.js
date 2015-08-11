function ShowBox(name) {
//  $print("started benchmark : " + name);
}

function AddResult(name, result) {
  $print(name + "\t" + result);
}

function AddError(name, error) {
//  $print("error: " + name + ": " + error.message);
}

function AddScore(score) {
// $print("score: " + score);
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
