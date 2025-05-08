import { useState } from 'react';
import './App.css';
import Modal from 'react-modal';

import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { BASE_URL } from './config';
import MediaList from './MediaList';

const MEDIA_TYPE = {
  IMAGE: 1,
  VIDEO: 2,
  FILE: 3,
  DEFAULT: 0,
  OSS: 99
};

function useMediaLoader() {
  const [mediaUrls, setMediaUrls] = useState([]);
  const [mediaType, setMediaType] = useState(MEDIA_TYPE.IMAGE);

  // 加载普通媒体
  const fetchMediaList = async (type) => {
    try {
      setMediaUrls([]);
      const requestData = {
        url: `${BASE_URL}/img/list`,
        method: 'POST',
        headers: {
          'Authorization': localStorage.getItem('token'),
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ type, mask: '' }),
        timestamp: new Date().toISOString()
      };
      const response = await fetch(requestData.url, {
        method: requestData.method,
        headers: requestData.headers,
        body: requestData.body
      });
      const data = await response.json();
      const uniqueUrls = Array.from(new Set((data.data||[]).map(item => item.url)));
      setMediaUrls(uniqueUrls);
      setMediaType(type);
      if(type === MEDIA_TYPE.VIDEO) toast.success('视频列表已加载');
    } catch (error) {
      toast.error('获取媒体失败');
    }
  };

  // 加载OSS媒体
  const fetchOssMediaList = async (ossType) => {
    try {
      setMediaUrls([]);
      const requestData = {
        url: `${BASE_URL}/img/oss/list`,
        method: 'POST',
        headers: {
          'Authorization': localStorage.getItem('token'),
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name: ossType }),
        timestamp: new Date().toISOString()
      };
      const response = await fetch(requestData.url, {
        method: requestData.method,
        headers: requestData.headers,
        body: requestData.body
      });
      const data = await response.json();
      const uniqueUrls = Array.from(new Set((data.data||[]).map(item => item.url)));
      setMediaUrls(uniqueUrls);
      setMediaType(MEDIA_TYPE.OSS);
      toast.success('OSS媒体列表已加载');
    } catch (error) {
      toast.error('获取OSS媒体失败');
    }
  };

  return { mediaUrls, setMediaUrls, mediaType, setMediaType, fetchMediaList, fetchOssMediaList };
}

export default function MediaGallery() {
  const { mediaUrls, setMediaUrls, mediaType, setMediaType, fetchMediaList, fetchOssMediaList } = useMediaLoader();
  const [selectedMedia, setSelectedMedia] = useState(null);
  const [expanded, setExpanded] = useState(true);

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
                setSelectedMedia(null);
                fetchMediaList(MEDIA_TYPE.FILE);
              }}
              style={{padding: '10px 30px', cursor: 'pointer', backgroundColor: mediaType === MEDIA_TYPE.FILE ? '#d9d9d9' : '#f0f2f5', color: '#000'}}
            >
              文件
            </div>
            <div 
              className="nav-subitem" 
              onClick={() => {
                setSelectedMedia(null);
                fetchOssMediaList('imgs');
              }}
              style={{padding: '10px 30px', cursor: 'pointer', backgroundColor: mediaType === MEDIA_TYPE.OSS && mediaUrls.length > 0 && mediaUrls[0].match(/(\.jpg|\.jpeg|\.png|\.gif)$/i) ? '#d9d9d9' : '#f0f2f5', color: '#000'}}
            >
              OSS图片
            </div>
            <div 
              className="nav-subitem" 
              onClick={() => {
                setSelectedMedia(null);
                fetchOssMediaList('vedios');
              }}
              style={{padding: '10px 30px', cursor: 'pointer', backgroundColor: mediaType === MEDIA_TYPE.OSS && mediaUrls.length > 0 && mediaUrls[0].match(/(\.mp4|\.avi|\.mov)$/i) ? '#d9d9d9' : '#f0f2f5', color: '#000'}}
            >
              OSS视频
            </div>
          </>
        )}
      </div>
      <div style={{display: 'flex', alignItems: 'center', gap: '12px', margin: '10px 0 0 10px'}}>
        <span style={{fontWeight: 500}}>布局切换：</span>
        <button className={'layout-btn'}>卡片</button>
      </div>
      <MediaList mediaUrls={mediaUrls} mediaType={mediaType} layout={'card'} onMediaSelect={setSelectedMedia} />
    </div>
  );
}
