<template>
	<div class="map-container">
		<div id="container"></div>
	</div>
</template>

<script setup>
import AMapLoader from '@amap/amap-jsapi-loader';
import { onMounted, onUnmounted, ref } from 'vue';

let map = null;

onMounted(() => {
	window._AMapSecurityConfig = {
		securityJsCode: 'b1e6cf67b54a313ecd56ee9489f9c954'
	};
	AMapLoader.load({
		key: 'ebd62b686942802c7218d0ee3aea566a', // 申请好的Web端开发者Key，首次调用 load 时必填
		version: '2.0', // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
		plugins: ['AMap.Scale'] //需要使用的的插件列表，如比例尺'AMap.Scale'，支持添加多个如：['...','...']
	})
		.then((AMap) => {
			map = new AMap.Map('container', {
				// 设置地图容器id
				viewMode: '3D', // 是否为3D地图模式
				zoom: 18, // 初始化地图级别
				center: [113.50932, 34.811214] // 初始化地图中心点位置
			});
		})
		.catch((e) => {
			console.log(e);
		});
});

onUnmounted(() => {
	map?.destroy();
});
</script>

<style lang="scss" scoped>
.map-container {
	padding: 20;
}

#container {
	background-color: slateblue;
	padding: 0px;
	margin: 0px;
	width: 100%;
	height: 600px;
}
</style>
