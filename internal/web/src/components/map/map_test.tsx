'use client';

import React, { useRef, useEffect, useState } from 'react';
import maplibregl from 'maplibre-gl';
import MaplibreGeocoder, {
    CarmenGeojsonFeature,
    MaplibreGeocoderApi,
    MaplibreGeocoderApiConfig,
} from '@maplibre/maplibre-gl-geocoder';
import 'maplibre-gl/dist/maplibre-gl.css';
import '@maplibre/maplibre-gl-geocoder/dist/maplibre-gl-geocoder.css';
import './custom.css';
import { createPortal } from 'react-dom';
import { Button } from '../shadcn-ui/button';
import { HousePlus } from 'lucide-react';

const API_KEY = process.env.NEXT_PUBLIC_MAP_STYLE_API_KEY;

function MapContainer() {
    const [, setClick] = useState(0);
    const mapContainer = useRef(null);
    const map = useRef<maplibregl.Map | null>(null);
    const contextMenu = useRef<maplibregl.Popup | null>(null);

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

        map.current.on('contextmenu', (e) => {
            if (!map.current) return;

            const lng = e.lngLat.lng;
            const lat = e.lngLat.lat;

            console.log(lng, lat);
            if (contextMenu.current) {
                contextMenu.current.setLngLat(e.lngLat).addTo(map.current);
            } else {
                contextMenu.current = new maplibregl.Popup({
                    closeButton: false,
                })
                    .setLngLat(e.lngLat)
                    .setHTML('<div id="popup"></div>')
                    .addTo(map.current);
            }

            setClick((p) => ++p);
        });

        const Geo = {
            forwardGeocode: async (config: MaplibreGeocoderApiConfig) => {
                const features: CarmenGeojsonFeature[] = [];
                try {
                    const request = `https://nominatim.openstreetmap.org/search?q=${config.query}&format=geojson&polygon_geojson=0&addressdetails=1`;
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

        const geocoder = new MaplibreGeocoder(Geo, {
            maplibregl: maplibregl,
            zoom: 14,
        });
        map.current.addControl(geocoder);
    }, [lng, lat, zoom]);

    return (
        <div className="relative w-full h-screen ">
            <div ref={mapContainer} className="absolute size-full" />
            {document
                ? document.getElementById('popup')
                    ? createPortal(
                          <Button className="px-2 -mt-1.5">
                              <HousePlus className="size-5 mr-1" /> Create
                          </Button>,
                          document.getElementById('popup') as HTMLElement
                      )
                    : null
                : null}
        </div>
    );
}

export default MapContainer;
