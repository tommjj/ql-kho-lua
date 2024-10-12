'use client';

import 'maplibre-gl/dist/maplibre-gl.css';
import '@maplibre/maplibre-gl-geocoder/dist/maplibre-gl-geocoder.css';
import './custom.css';

import React, {
    DetailedHTMLProps,
    HTMLAttributes,
    useCallback,
    useRef,
    useImperativeHandle,
} from 'react';

import Map, {
    MapRef,
    NavigationControl,
    FullscreenControl,
    ScaleControl,
    GeolocateControl,
} from 'react-map-gl/maplibre';
import MaplibreGeocoder from '@maplibre/maplibre-gl-geocoder';
import maplibregl from 'maplibre-gl';
import { geocoder } from '@/lib/map/geocoder';
import { cn } from '@/lib/utils';

const API_KEY = process.env.NEXT_PUBLIC_MAP_STYLE_API_KEY;

type PropsType = DetailedHTMLProps<
    HTMLAttributes<HTMLDivElement>,
    HTMLDivElement
>;
export type MapRefType = {
    flyTo(options: maplibregl.FlyToOptions): void;
};

const MapContainer = React.forwardRef<MapRefType, PropsType>(
    function MapContainer({ children, className, ...props }: PropsType, ref) {
        const mapRef = useRef<MapRef | undefined>(undefined);

        useImperativeHandle(ref, () => ({
            flyTo(options) {
                const map = mapRef.current;
                if (!map) {
                    return;
                }

                map.flyTo(options);
            },
        }));

        const initMaplibreGeocoder = useCallback(() => {
            const map = mapRef.current;
            if (!map) {
                return;
            }

            map.flyTo;

            const geo = new MaplibreGeocoder(geocoder, {
                maplibregl: maplibregl,
                zoom: 14,
            });
            map.addControl(geo, 'top-right');
        }, []);

        return (
            <div
                className={cn('relative w-full h-screen', className)}
                {...props}
            >
                <Map
                    onLoad={initMaplibreGeocoder}
                    // eslint-disable-next-line @typescript-eslint/no-explicit-any
                    ref={mapRef as any}
                    initialViewState={{
                        longitude: -100,
                        latitude: 40,
                        zoom: 3.5,
                    }}
                    mapStyle={`https://api.maptiler.com/maps/streets-v2/style.json?key=${API_KEY}`}
                >
                    <GeolocateControl position="top-left" />
                    <FullscreenControl position="top-left" />
                    <NavigationControl position="top-left" />
                    <ScaleControl />
                    {children}
                </Map>
            </div>
        );
    }
);

export default MapContainer;
