import { useState } from 'react';
import './App.css';
import Modal from 'react-modal';

import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { BASE_URL } from './config';

const MEDIA_TYPE = {
  IMAGE: 1,
  VIDEO: 2,
  FILE: 3,
  DEFAULT: 0
};

export default function MediaGallery() {
  const [mediaUrls, setMediaUrls] = useState([]);
  const [selectedMedia, setSelectedMedia] = useState(null);
  const [expanded, setExpanded] = useState(true);

  const [mediaType, setMediaType] = useState(MEDIA_TYPE.IMAGE);

  const fetchMediaList = async (type) => {
    try {
      const requestData = {
        url: `${BASE_URL}/img/list`,
        method: 'POST',
        headers: {
          'Authorization': localStorage.getItem('token'),
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          type: type,
          mask: ''
        }),
        timestamp: new Date().toISOString()
      };
      

      const response = await fetch(requestData.url, {
        method: requestData.method,
        headers: requestData.headers,
        body: requestData.body
      });
      const data = await response.json();
      setMediaUrls(data.data.map(item => item.url));
      if(type === MEDIA_TYPE.VIDEO) {
        toast.success('视频列表已加载');
      }
    } catch (error) {
      toast.error('获取媒体失败');
    }
  };

  return (
    <div className="media-gallery-container" style={{marginLeft: '200px'}}>
      <ToastContainer 
        position="top-right"
        autoClose={3000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
      />
      
      <div className="sidebar" style={{width: '200px', height: '100%', position: 'fixed', left: 0, top: 0, backgroundColor: '#f0f2f5', padding: '0'}}>
        <div 
          className="nav-item" 
          onClick={() => setExpanded(!expanded)}
          style={{padding: '10px 20px', cursor: 'pointer', backgroundColor: expanded ? '#d9d9d9' : '#f0f2f5', color: '#000'}}
        >
          文件资源
        </div>
        {expanded && (
          <>
            <div 
              className="nav-subitem" 
              onClick={() => {
                setMediaType(MEDIA_TYPE.IMAGE);
                setMediaUrls([]);
                setSelectedMedia(null);
                fetchMediaList(MEDIA_TYPE.IMAGE);
              }}
              style={{padding: '10px 30px', cursor: 'pointer', backgroundColor: mediaType === MEDIA_TYPE.IMAGE ? '#d9d9d9' : '#f0f2f5', color: '#000'}}
            >
              图片
            </div>
            <div 
              className="nav-subitem" 
              onClick={() => {
                setMediaType(MEDIA_TYPE.VIDEO);
                setMediaUrls([]);
                setSelectedMedia(null);
                fetchMediaList(MEDIA_TYPE.VIDEO);
              }}
              style={{padding: '10px 30px', cursor: 'pointer', backgroundColor: mediaType === MEDIA_TYPE.VIDEO ? '#d9d9d9' : '#f0f2f5', color: '#000'}}
            >
              视频
            </div>
            <div 
              className="nav-subitem" 
              onClick={() => {
                setMediaType(MEDIA_TYPE.FILE);
                setMediaUrls([]);
                setSelectedMedia(null);
                fetchMediaList(MEDIA_TYPE.FILE);
              }}
              style={{padding: '10px 30px', cursor: 'pointer', backgroundColor: mediaType === MEDIA_TYPE.FILE ? '#d9d9d9' : '#f0f2f5', color: '#000'}}
            >
              文件
            </div>
          </>
        )}
      </div>
      
      <div className="media-list" style={{display: 'flex', flexWrap: 'wrap', gap: '10px', padding: '10px'}}>
        {mediaUrls.map((url, index) => (
          <div key={index} className="media-list-item">
            <div className="media-content">
              {mediaType === MEDIA_TYPE.IMAGE ? (
                <img 
                  src={url} 
                  alt={`图片${index}`} 
                  className="list-image" 
                  onClick={() => window.open(url, '_blank')}
                  style={{width: '150px', height: '150px', objectFit: 'cover', margin: '5px', cursor: 'pointer'}}
                />
              ) : mediaType === MEDIA_TYPE.VIDEO ? (
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
                  onDoubleClick={() => setSelectedMedia(url)}
                  style={{width: '100%', height: 'auto'}}
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
            <div className="media-actions" style={{display: 'flex', justifyContent: 'flex-end', alignItems: 'center', gap: '8px'}}>

              
            </div>
          </div>
        ))}
      </div>
      
      <Modal 
        isOpen={!!selectedMedia}
        onRequestClose={() => setSelectedMedia(null)}
        contentLabel="媒体预览"
        className="media-modal"
        overlayClassName="modal-overlay"
        style={{
          overlay: {
            zIndex: 9999,
            backgroundColor: 'rgba(0, 0, 0, 0.75)'
          },
          content: {
            top: '50%',
            left: '50%',
            right: 'auto',
            bottom: 'auto',
            marginRight: '-50%',
            transform: 'translate(-50%, -50%)',
            padding: '20px',
            maxWidth: '90vw',
            maxHeight: '90vh',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            flexDirection: 'column',
            width: 'auto',
            height: 'auto'
          }
        }}
      >
        <button 
          onClick={() => setSelectedMedia(null)} 
          style={{
            position: 'absolute', 
            top: '10px', 
            right: '10px',
            background: 'none',
            border: 'none',
            fontSize: '20px',
            cursor: 'pointer',
            color: '#333'
          }}
        >
          ×
        </button>
        {selectedMedia && (
          mediaType === MEDIA_TYPE.IMAGE ? (
            <img 
              src={selectedMedia} 
              alt="预览" 
              className="modal-image"
              style={{maxWidth: '90%', maxHeight: '90%', objectFit: 'contain'}}
            />
          ) : (
            <video 
              src={selectedMedia} 
              controls
              className="modal-video"
              autoPlay
              onClick={(e) => e.stopPropagation()}
            />
          )
        )}
      </Modal>
    </div>
  );
}