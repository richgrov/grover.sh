import * as THREE from "three";

const epsilon = (value: number) => (Math.abs(value) < 1e-10 ? 0 : value);

function matrixToCss(matrix: THREE.Matrix4, multipliers: number[]) {
  let matrix3d = "matrix3d(";
  for (let i = 0; i < 16; i++) {
    matrix3d +=
      epsilon(multipliers[i]! * matrix.elements[i]!) + (i !== 15 ? "," : ")");
  }
  return matrix3d;
}

function cameraToCss(matrix: THREE.Matrix4) {
  return matrixToCss(
    matrix,
    [1, -1, 1, 1, 1, -1, 1, 1, 1, -1, 1, 1, 1, -1, 1, 1]
  );
}

function objectToCss(matrix: THREE.Matrix4) {
  return (
    "translate(-50%,-50%)" +
    matrixToCss(matrix, [1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1, 1, 1, 1])
  );
}

const plane = new THREE.PlaneGeometry(1, 1);

function createPlane(
  width: number,
  height: number,
  color: number,
  position: THREE.Vector3 = new THREE.Vector3(0, 0, 0),
  rotation: THREE.Euler = new THREE.Euler(0, 0, 0)
) {
  const obj = new THREE.Mesh(plane, new THREE.MeshBasicMaterial({ color }));
  obj.position.copy(position);
  obj.rotation.copy(rotation);
  obj.scale.set(width, height, 1);
  return obj;
}

export function start() {
  document.body.requestPointerLock();
  const maxScroll = document.body.scrollHeight;
  const scrollY = window.scrollY / document.body.scrollHeight;
  window.scrollTo(0, 0);
  const aspect = window.innerWidth / window.innerHeight;
  const scene = new THREE.Scene();
  const camera = new THREE.PerspectiveCamera(75, aspect, 0.1, 1000);
  camera.position.z = 1;

  const renderer = new THREE.WebGLRenderer();
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.setPixelRatio(window.devicePixelRatio);

  const planeHeight = 2 * Math.tan(0.5 * ((camera.fov * Math.PI) / 180));
  const planeWidth = planeHeight * aspect;

  scene.add(createPlane(planeWidth, planeHeight, 0xfefefa));
  scene.add(
    createPlane(
      planeWidth,
      planeHeight,
      0xff0000,
      new THREE.Vector3(-planeWidth / 2, 0, planeWidth / 2),
      new THREE.Euler(0, Math.PI / 2, 0)
    )
  );
  scene.add(
    createPlane(
      planeWidth,
      planeHeight,
      0xff0000,
      new THREE.Vector3(planeWidth / 2, 0, planeWidth / 2),
      new THREE.Euler(0, -Math.PI / 2, 0)
    )
  );

  const keys: Record<string, boolean> = {};
  let mouseX = 0;
  let mouseY = 0;
  let isPointerLocked = false;

  const projection = document.getElementById("projection") as HTMLDivElement;
  const viewEl = document.getElementById("view") as HTMLDivElement;
  const objectEl = document.getElementById("model") as HTMLDivElement;

  let time = performance.now();

  function loop() {
    renderer.render(scene, camera);

    const now = performance.now();
    const delta = (now - time) / 1000;
    time = now;

    window.requestAnimationFrame(loop);

    const moveSpeed = 1;

    const forward = new THREE.Vector3(
      Math.sin(camera.rotation.y),
      0,
      Math.cos(camera.rotation.y)
    );
    const right = new THREE.Vector3(
      Math.cos(camera.rotation.y),
      0,
      -Math.sin(camera.rotation.y)
    );

    if (keys["w"] || keys["W"]) {
      camera.position.addScaledVector(forward, -moveSpeed * delta);
    }
    if (keys["s"] || keys["S"]) {
      camera.position.addScaledVector(forward, moveSpeed * delta);
    }
    if (keys["a"] || keys["A"]) {
      camera.position.addScaledVector(right, -moveSpeed * delta);
    }
    if (keys["d"] || keys["D"]) {
      camera.position.addScaledVector(right, moveSpeed * delta);
    }

    camera.rotation.y = mouseX;
    camera.rotation.x = mouseY;
    camera.rotation.order = "YXZ";

    camera.updateMatrixWorld();

    const widthHalf = window.innerWidth / 2;
    const heightHalf = window.innerHeight / 2;
    const fov = camera.projectionMatrix.elements[5] * heightHalf;
    const cameraMatrix = cameraToCss(camera.matrixWorldInverse);
    const cameraTransform = `translateZ(${fov}px)`;

    projection.style.perspective = `${fov}px`;
    projection.style.width = window.innerWidth + "px";
    projection.style.height = window.innerHeight + "px";

    viewEl.style.position = "fixed";
    viewEl.style.top = "0";
    viewEl.style.left = "0";
    viewEl.style.width = window.innerWidth + "px";
    viewEl.style.height = window.innerHeight + "px";
    viewEl.style.transformStyle = "preserve-3d";
    viewEl.style.pointerEvents = "none";
    viewEl.style.transform = `${cameraTransform}${cameraMatrix}translate(${widthHalf}px,${heightHalf}px)`;

    objectEl.style.position = "absolute";
    objectEl.style.pointerEvents = "auto";
    const domScale = planeWidth / window.innerWidth;
    const planeMatrix = new THREE.Matrix4().compose(
      new THREE.Vector3(
        0,
        maxScroll * domScale * (scrollY - 0.5) + planeHeight / 2,
        0
      ),
      new THREE.Quaternion(),
      new THREE.Vector3().setScalar(domScale)
    );

    objectEl.style.transform = objectToCss(planeMatrix);
  }

  window.requestAnimationFrame(loop);

  window.onkeydown = (event) => {
    keys[event.key] = true;
  };

  window.onkeyup = (event) => {
    keys[event.key] = false;
  };

  window.onmousemove = (event) => {
    if (isPointerLocked) {
      mouseX -= event.movementX * 0.002;
      mouseY -= event.movementY * 0.002;
      mouseY = Math.max(-Math.PI / 2, Math.min(Math.PI / 2, mouseY));
    }
  };

  window.onclick = () => {
    if (!isPointerLocked) {
      document.body.requestPointerLock();
    }
  };

  document.addEventListener("pointerlockchange", () => {
    isPointerLocked = document.pointerLockElement === document.body;
  });

  document.body.appendChild(renderer.domElement);
  document.body.classList.remove("with-margin");
  document.body.classList.add("game");
  projection.style.transformOrigin = "0 0";
}
