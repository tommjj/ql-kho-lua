'use client';

import { useCallback, useState } from 'react';
import MapContainer from '../map/map';
import { Button } from '../shadcn-ui/button';
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from '../shadcn-ui/dialog';
import { Input } from '../shadcn-ui/input';
import { Label } from '../shadcn-ui/label';
import UploadImageSelect from '../ui/file-input';
import { MapLayerMouseEvent } from 'maplibre-gl';
import { Marker } from 'react-map-gl/maplibre';
import { AlertCircle, FilePenLine } from 'lucide-react';
import Pin from '../map/pin';
import { LngLat } from '@/types/data-types';
import { isLocation, parseLocation } from '@/lib/validator/location';
import { cn } from '@/lib/utils';
import { useRouter } from 'next/navigation';
import {
    updateWarehouse,
    UpdateWarehouseSchema,
} from '@/lib/services/warehouse.service';
import { useSession } from '../session-context';
import { Cross2Icon } from '@radix-ui/react-icons';
import { Warehouse } from '@/lib/zod.schema';

export function UpdateWarehouse({
    warehouse,
    onOpenChange,
}: {
    warehouse: Warehouse;
    onOpenChange?: (open: boolean) => void;
}) {
    const { refresh } = useRouter();
    const user = useSession();
    const [isOpen, setIsOpen] = useState(false);

    const [location, setLocation] = useState<LngLat | null>({
        lat: warehouse.location[0],
        lng: warehouse.location[1],
    });
    const [locationString, setLocationString] = useState<string>(
        `${warehouse.location[0]},${warehouse.location[1]}`
    );
    const [name, setName] = useState(warehouse.name);
    const [capacity, setCapacity] = useState<null | number>(warehouse.capacity);
    const [image, setImage] = useState<string | undefined>(undefined);

    const [error, setError] = useState<string | null>();

    const handleSelectLocation = useCallback((e: MapLayerMouseEvent) => {
        setLocationString(`${e.lngLat.lat},${e.lngLat.lng}`);

        setLocation({
            lat: e.lngLat.lat,
            lng: e.lngLat.lng,
        });
    }, []);

    const handleCreate = useCallback(
        (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
            e.preventDefault();

            const createData = {
                id: warehouse.id,
                name: name,
                location: [location?.lat, location?.lng],
                capacity: capacity,
                image: image,
            };

            const parse = UpdateWarehouseSchema.safeParse(createData);

            if (!parse.success) return;

            (async () => {
                const [, err] = await updateWarehouse(user.token, parse.data);
                if (!err) {
                    refresh();
                    setIsOpen(false);
                    return;
                }
                if (!(err instanceof Response)) return;

                switch (err.status) {
                    case 409:
                        setError('warehouse name is exist');
                        break;
                    case 400:
                        const data = await err.json();
                        setError(data.messages[0] as string);
                        break;
                }
            })();

            setError(null);
        },
        [
            capacity,
            image,
            location?.lat,
            location?.lng,
            name,
            refresh,
            warehouse.id,
            user.token,
        ]
    );

    return (
        <Dialog open={isOpen} onOpenChange={onOpenChange}>
            <DialogTrigger asChild onClick={() => setIsOpen(true)}>
                <Button
                    className="flex items-center justify-start w-full"
                    variant={'ghost'}
                >
                    <FilePenLine className="size-5 mr-2 " />
                    <span>Edit</span>
                </Button>
            </DialogTrigger>
            <DialogContent className="flex min-w-[90vw] h-[85vh] bg-white">
                <DialogClose
                    onClick={() => setIsOpen(false)}
                    className="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground"
                >
                    <Cross2Icon className="h-4 w-4" />
                    <span className="sr-only">Close</span>
                </DialogClose>
                <div className="w-[700px] h-full rounded overflow-hidden ">
                    <MapContainer
                        initialViewState={{
                            longitude: location?.lng,
                            latitude: location?.lat,
                            zoom: 3.5,
                        }}
                        onClick={handleSelectLocation}
                        minZoom={2}
                    >
                        {location ? (
                            <Marker
                                longitude={location.lng}
                                latitude={location.lat}
                                anchor="bottom"
                            >
                                <Pin />
                            </Marker>
                        ) : null}
                    </MapContainer>
                </div>
                <div className="flex flex-col flex-grow">
                    <DialogHeader>
                        <DialogTitle className="text-2xl">
                            Edit warehouse
                        </DialogTitle>
                        <DialogDescription>
                            Edit info warehouse here. Click Save when you are
                            done.
                        </DialogDescription>
                        {error && (
                            <div
                                className="flex items-center text-destructive"
                                id="file-upload-error"
                                role="alert"
                            >
                                <AlertCircle className="h-4 w-4 mr-2" />
                                <span className="text-sm">{error}</span>
                            </div>
                        )}
                    </DialogHeader>
                    <div className="flex-grow">
                        <div className=" grid gap-4 py-4">
                            <div className=" items-center gap-4">
                                <Label htmlFor="name" className="text-right">
                                    Name
                                </Label>
                                <Input
                                    id="name"
                                    onChange={(e) => setName(e.target.value)}
                                    value={name}
                                    className={cn('col-span-3', {
                                        'focus-visible:ring-red-700':
                                            name.length <= 3,
                                    })}
                                />
                            </div>
                            <div className=" gap-4">
                                <Label
                                    htmlFor="location"
                                    className="text-right"
                                >
                                    Location
                                </Label>
                                <Input
                                    id="location"
                                    value={locationString}
                                    onChange={(e) => {
                                        const value = e.target.value;

                                        setLocationString(value);
                                        const [l, ok] = parseLocation(value);
                                        setLocation(ok ? l : null);
                                    }}
                                    className={cn('col-span-3', {
                                        'focus-visible:ring-red-700':
                                            !isLocation(locationString),
                                    })}
                                />
                            </div>
                            <div className=" gap-4">
                                <Label
                                    htmlFor="capacity"
                                    className="text-right"
                                >
                                    Capacity
                                </Label>
                                <Input
                                    type="number"
                                    id="capacity"
                                    onChange={(e) =>
                                        setCapacity(Number(e.target.value))
                                    }
                                    value={capacity ? capacity : ''}
                                    className={cn('col-span-3', {
                                        'focus-visible:ring-red-700':
                                            capacity && capacity < 1,
                                    })}
                                />
                            </div>
                        </div>

                        <div>
                            <UploadImageSelect
                                defaultImage={warehouse.image}
                                className="w-full"
                                onUploaded={(filename) =>
                                    filename
                                        ? setImage(filename)
                                        : setImage(undefined)
                                }
                            />
                        </div>
                    </div>
                    <DialogFooter>
                        <Button onClick={handleCreate} type="submit" asChild>
                            <DialogClose>Save</DialogClose>
                        </Button>
                    </DialogFooter>
                </div>
            </DialogContent>
        </Dialog>
    );
}
