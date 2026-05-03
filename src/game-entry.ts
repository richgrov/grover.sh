import * as game from "./game";

var gameInitialized = false;

document.onkeydown = (e) => {
  switch (e.key) {
    case "w":
    case "a":
    case "s":
    case "d":
      e.preventDefault();
      break;
    default:
      return;
  }

  if (gameInitialized) {
    return;
  }

  gameInitialized = true;
  game.start();
};
