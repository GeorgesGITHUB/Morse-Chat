
function playMorse(morseString) {
    const context = new (window.AudioContext || window.webkitAudioContext)();
    const gain = context.createGain();
    let oscillator = context.createOscillator();

    const dot = 1.2 / 15;
    let t = context.currentTime;

    oscillator.type = "sine";
    oscillator.frequency.value = 600;

    gain.gain.setValueAtTime(0, t);

    morseString.split("").forEach(function(letter) {
        switch(letter) {
            case ".":
                gain.gain.setValueAtTime(1, t);
                t += dot;
                gain.gain.setValueAtTime(0, t);
                t += dot;
                break;
            case "-":
                gain.gain.setValueAtTime(1, t);
                t += 3 * dot;
                gain.gain.setValueAtTime(0, t);
                t += dot;
                break;
            case " ":
                t += 7 * dot;
                break;
        }
    });

    oscillator.connect(gain);
    gain.connect(context.destination);

    oscillator.start();

    return false;
}

export default playMorse