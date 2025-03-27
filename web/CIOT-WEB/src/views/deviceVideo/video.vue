<template>
  <div class="video-player-container">
    <!-- 视频播放器 -->
    <video ref="videoPlayer" class="video-js vjs-default-skin" controls preload="auto" width="640" height="360"></video>

    <!-- 设备信息 -->
    <div class="device-info" v-if="deviceVideo">
      <h3>设备视频信息</h3>
      <p><strong>设备名称:</strong> {{ deviceVideo.device_name }}</p>
      <p><strong>设备ID:</strong> {{ deviceVideo.device_id }}</p>
      <p><strong>视频状态:</strong> {{ deviceVideo.status === 1 ? '在线' : '离线' }}</p>
      <p><strong>视频流类型:</strong> {{ deviceVideo.stream_type }}</p>
      <p><strong>视频流地址:</strong> {{ deviceVideo.video_url }}</p>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading">加载中...</div>

    <!-- 错误提示 -->
    <div v-if="error" class="error">加载失败: {{ error }}</div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';
import videojs from 'video.js';
import 'video.js/dist/video-js.css';
import { getDeviceVideoDetail } from '@/api/deviceVideo'; // 根据你的路径引入
import { useRouter, useRoute } from 'vue-router'

const route = useRoute()
const router = useRouter()
const videoPlayer = ref(null); // 视频播放器 DOM 元素
const deviceVideo = ref(null); // 设备视频信息
const loading = ref(false); // 加载状态
const error = ref(''); // 错误信息
let player = null; // video.js 播放器实例

// 获取设备视频信息
const fetchDeviceVideo = async () => {
  loading.value = true;
  error.value = '';

  try {
    const response = await getDeviceVideoDetail(route.params.id);
    deviceVideo.value = response.data;

    // 初始化视频播放器
    if (deviceVideo.value.stream_type.toLowerCase() === 'hls' && deviceVideo.value.video_url) {
      initializeHLSPlayer();
    } else {
      error.value = '不支持的视频流类型';
    }
  } catch (err) {
    error.value = err.message;
  } finally {
    loading.value = false;
  }
};

// 初始化 HLS 视频播放器
const initializeHLSPlayer = () => {
  if (deviceVideo.value.enable_video && deviceVideo.value.video_url) {
    player = videojs(videoPlayer.value, {
      techOrder: ['html5'], // 使用 HTML5 技术播放
      sources: [{
        src: deviceVideo.value.video_url, // 视频流地址
        type: 'application/x-mpegURL', // HLS 类型
      }],
      controls: true, // 显示控制条
      autoplay: true, // 自动播放
      preload: 'auto', // 预加载
    });

    // 监听播放器错误
    player.on('error', () => {
      console.error('视频播放错误，请检查视频流地址或网络连接');
    });
  } else {
    console.warn('视频功能未启用或视频流地址无效');
  }
};

// 组件挂载时获取数据
onMounted(() => {
  fetchDeviceVideo();
});

// 组件销毁时销毁播放器
onBeforeUnmount(() => {
  if (player) {
    player.dispose();
  }
});
</script>

<style scoped>
.video-player-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.device-info {
  background-color: #fff;
  padding: 15px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 640px;
}

.device-info h3 {
  margin-top: 0;
  color: #333;
}

.device-info p {
  margin: 5px 0;
  color: #666;
}

.video-js {
  background-color: #000;
  border-radius: 8px;
  overflow: hidden;
}

.loading,
.error {
  color: #666;
  font-size: 16px;
  margin-top: 20px;
}

.error {
  color: red;
}
</style>
