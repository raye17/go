import React from 'react';

const MEDIA_TYPE = {
  IMAGE: 1,
  VIDEO: 2,
  FILE: 3,
  DEFAULT: 0,
  OSS: 99
};

function getMediaType(url) {
  if (/\.(jpg|jpeg|png|gif)$/i.test(url)) return MEDIA_TYPE.IMAGE;
  if (/\.(mp4|avi|mov)$/i.test(url)) return MEDIA_TYPE.VIDEO;
  return MEDIA_TYPE.FILE;
}

export default function MediaList({ mediaUrls, mediaType, layout, onMediaSelect }) {
  return (
    <div className={layout === 'grid' ? 'media-list grid-layout' : 'media-list card-layout'} style={{gap: '10px', padding: '10px'}}>
      {mediaUrls.map((url, index) => {
        const type = mediaType === MEDIA_TYPE.OSS ? getMediaType(url) : mediaType;
        return (
          <div key={index} className="media-list-item">
            <div className="media-content">
              {type === MEDIA_TYPE.IMAGE ? (
                <img 
                  src={url} 
                  alt={`图片${index}`} 
                  className="list-image"
                  onClick={() => window.open(url, '_blank')}
                />
              ) : type === MEDIA_TYPE.VIDEO ? (
                <video 
                  src={url} 
                  className="list-video"
                  onClick={(e) => {
                    e.stopPropagation();
                    const video = e.target;
                    if(video.paused) {
                      video.play();
                    } else {
                      video.pause();
                    }
                  }}
                  onDoubleClick={() => onMediaSelect(url)}
                  controls
                />
              ) : (
                <div className="file-list-item">
                  <div className="file-icon">📄</div>
                  <div className="file-info">
                    <div className="file-name">{url.split('/').pop()}</div>
                    <div className="file-type">文件</div>
                  </div>
                </div>
              )}
            </div>
            <div className="media-actions" style={{display: 'flex', justifyContent: 'flex-end', alignItems: 'center', gap: '8px'}}></div>
          </div>
        );
      })}
    </div>
  );
}

