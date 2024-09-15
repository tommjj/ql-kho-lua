'use client';

import React, { useRef, useEffect } from 'react';
import maplibregl from 'maplibre-gl';
import MaplibreGeocoder, {
    CarmenGeojsonFeature,
    MaplibreGeocoderApi,
    MaplibreGeocoderApiConfig,
} from '@maplibre/maplibre-gl-geocoder';
import 'maplibre-gl/dist/maplibre-gl.css';
import '@maplibre/maplibre-gl-geocoder/dist/maplibre-gl-geocoder.css';

const API_KEY = process.env.NEXT_PUBLIC_MAP_STYLE_API_KEY;

function MapContainer() {
    const mapContainer = useRef(null);
    const map = useRef<maplibregl.Map | null>(null);
    const marker = useRef<maplibregl.Marker | null>(null);

    const lng = 105.09185355014682;
    const lat = 9.982425277100205;
    const zoom = 14;

    useEffect(() => {
        if (map.current) return; // stops map from intializing more than once
        if (!mapContainer.current) return;

        map.current = new maplibregl.Map({
            container: mapContainer.current,
            style: `https://api.maptiler.com/maps/streets-v2/style.json?key=${API_KEY}`,
            center: [lng, lat],
            zoom: zoom,
        });

        map.current.addControl(new maplibregl.NavigationControl(), 'top-left');

        map.current.on('click', (e) => {
            if (!map.current) return;

            const lng = e.lngLat.lng;
            const lat = e.lngLat.lat;

            console.log(lng, lat);

            if (marker.current) {
                marker.current.setLngLat([lng, lat]);
            } else {
                marker.current = new maplibregl.Marker({
                    color: '#FF0000',
                });

                marker.current.setLngLat([lng, lat]).addTo(map.current);
            }
        });

        const Geo = {
            forwardGeocode: async (config: MaplibreGeocoderApiConfig) => {
                const features: CarmenGeojsonFeature[] = [];
                try {
                    const request = `https://nominatim.openstreetmap.org/search?q=${config.query}&format=geojson&polygon_geojson=1&addressdetails=1`;
                    const response = await fetch(request);
                    const geojson = await response.json();
                    for (const feature of geojson.features) {
                        const center = [
                            feature.bbox[0] +
                                (feature.bbox[2] - feature.bbox[0]) / 2,
                            feature.bbox[1] +
                                (feature.bbox[3] - feature.bbox[1]) / 2,
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

        const geocoder = new MaplibreGeocoder(Geo, {
            maplibregl: maplibregl,
            zoom: 14,
        });
        map.current.addControl(geocoder);
    }, [lng, lat, zoom]);

    return (
        <div className="relative w-screen h-screen">
            <div ref={mapContainer} className="absolute size-full" />
        </div>
    );
}

export default MapContainer;
