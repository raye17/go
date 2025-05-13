import React, { useRef, useEffect, useState } from 'react';
import './Cube3D.css';

export default function Cube3D({ images = [] }) {
  const [rotate, setRotate] = useState({ x: 0, y: 0 });
  const dragging = useRef(false);
  const lastPos = useRef({ x: 0, y: 0 });

  const handleMouseDown = (e) => {
    dragging.current = true;
    lastPos.current = { x: e.clientX, y: e.clientY };
  };

  const handleMouseMove = (e) => {
    if (!dragging.current) return;
    const dx = e.clientX - lastPos.current.x;
    const dy = e.clientY - lastPos.current.y;
    setRotate((prev) => ({ x: prev.x + dy, y: prev.y + dx }));
    lastPos.current = { x: e.clientX, y: e.clientY };
  };

  const handleMouseUp = () => {
    dragging.current = false;
  };

  useEffect(() => {
    window.addEventListener('mousemove', handleMouseMove);
    window.addEventListener('mouseup', handleMouseUp);
    return () => {
      window.removeEventListener('mousemove', handleMouseMove);
      window.removeEventListener('mouseup', handleMouseUp);
    };
  });

  // 立方体六面图片，若不足6张补空
  const faces = [];
  for (let i = 0; i < 6; i++) {
    faces.push(images[i] || '');
  }

  return (
    <div className="cube3d-scene">
      <div
        className="cube3d-cube"
        style={{
          transform: `rotateX(${rotate.x}deg) rotateY(${rotate.y}deg)`
        }}
        onMouseDown={handleMouseDown}
      >
        <div className="cube3d-face cube3d-face-front">
          {faces[0] && <img src={faces[0]} alt="front" />}
        </div>
        <div className="cube3d-face cube3d-face-back">
          {faces[1] && <img src={faces[1]} alt="back" />}
        </div>
        <div className="cube3d-face cube3d-face-right">
          {faces[2] && <img src={faces[2]} alt="right" />}
        </div>
        <div className="cube3d-face cube3d-face-left">
          {faces[3] && <img src={faces[3]} alt="left" />}
        </div>
        <div className="cube3d-face cube3d-face-top">
          {faces[4] && <img src={faces[4]} alt="top" />}
        </div>
        <div className="cube3d-face cube3d-face-bottom">
          {faces[5] && <img src={faces[5]} alt="bottom" />}
        </div>
      </div>
      <div className="cube3d-tip">按住鼠标拖动旋转立方体</div>
    </div>
  );
}