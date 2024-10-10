import {
    CarmenGeojsonFeature,
    MaplibreGeocoderApi,
    MaplibreGeocoderApiConfig,
} from '@maplibre/maplibre-gl-geocoder';

export const geocoder = {
    forwardGeocode: async (config: MaplibreGeocoderApiConfig) => {
        const features: CarmenGeojsonFeature[] = [];
        try {
            const request = `https://nominatim.openstreetmap.org/search?q=${config.query}&format=geojson&polygon_geojson=0&addressdetails=1`;
            const response = await fetch(request);
            const geojson = await response.json();
            for (const feature of geojson.features) {
                const center = [
                    feature.bbox[0] + (feature.bbox[2] - feature.bbox[0]) / 2,
                    feature.bbox[1] + (feature.bbox[3] - feature.bbox[1]) / 2,
                ];
                const point: CarmenGeojsonFeature = {
                    id: feature.properties.place_id,
                    type: 'Feature',
                    geometry: {
                        type: 'Point',
                        coordinates: center,
                    },
                    place_name: feature.properties.display_name,
                    properties: feature.properties,
                    text: feature.properties.display_name,
                    place_type: ['place'],
                    bbox: feature.bbox,
                };
                features.push(point);
            }
        } catch (e) {
            console.error(`Failed to forwardGeocode with error: ${e}`);
        }

        return {
            type: 'FeatureCollection',
            features: features,
        };
    },
} as MaplibreGeocoderApi;
