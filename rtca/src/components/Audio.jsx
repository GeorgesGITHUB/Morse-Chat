
function test(morseString) {
    const ctx = new AudioContext();
    const dot = 1.2 / 15;
    let t = ctx.currentTime;

    let oscillator = ctx.createOscillator();
    oscillator.type = "sine";
    oscillator.frequency.value = 600;

    const gainNode = ctx.createGain();
    gainNode.gain.setValueAtTime(0, t);

    morseString.split("").forEach(function(letter) {
        switch(letter) {
            case ".":
                gainNode.gain.setValueAtTime(1, t);
                t += dot;
                gainNode.gain.setValueAtTime(0, t);
                t += dot;
                break;
            case "-":
                gainNode.gain.setValueAtTime(1, t);
                t += 3 * dot;
                gainNode.gain.setValueAtTime(0, t);
                t += dot;
                break;
            case " ":
                t += 7 * dot;
                break;
        }
    });

    oscillator.connect(gainNode);
    gainNode.connect(ctx.destination);

    oscillator.start();

    return false;
}

export default test