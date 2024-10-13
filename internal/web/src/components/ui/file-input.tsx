'use client';

import Image from 'next/image';
import { useState, useRef } from 'react';
import { Input } from '@/components/shadcn-ui/input';
import { Label } from '@/components/shadcn-ui/label';
import { AlertCircle, Upload } from 'lucide-react';
import { useSession } from '../session-context';
import { uploadImageFile } from '@/lib/services/upload.service';
import { cn } from '@/lib/utils';

type Props = {
    className?: string;
    onUploaded?: (filename: string) => void;
};

export default function UploadImageSelect({ className, onUploaded }: Props) {
    const [uploading, setUploading] = useState(false);
    const userSession = useSession();
    const [selectedImage, setSelectedImage] = useState<string | null>(null);
    const [error, setError] = useState<string | null>(null);
    const fileInputRef = useRef<HTMLInputElement>(null);

    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        setError(null);

        if (file) {
            if (!file.type.startsWith('image/')) return;
            setUploading(true);
            const formData = new FormData();
            formData.append('file', file);

            setSelectedImage(null);

            uploadImageFile(userSession.token, file).then(
                async ([res, err]) => {
                    if (!err) {
                        onUploaded && onUploaded(res.data.filename);

                        setSelectedImage(res.data.filename);
                    } else {
                        onUploaded && onUploaded('');
                        setError('Please select an image file.');
                        if (fileInputRef.current) {
                            fileInputRef.current.value = '';
                        }
                    }
                    setUploading(false);
                }
            );
        } else {
            setSelectedImage(null);
        }
    };

    return (
        <div
            className={cn(
                'max-w-4xl w-full max-h-28 mx-auto bg-card rounded-lg',
                className
            )}
        >
            <div className="flex flex-col md:flex-row md:space-x-6 space-y-4 md:space-y-0">
                <div className="md:w-1/2 flex items-center justify-center bg-muted rounded-lg p-4">
                    {selectedImage ? (
                        // eslint-disable-next-line @next/next/no-img-element
                        <Image
                            width={180}
                            height={120}
                            src={`/temp/${selectedImage}`}
                            alt="Selected image preview"
                            className="max-w-full max-h-28 h-28 object-contain rounded-md"
                        />
                    ) : (
                        <div className="text-center text-muted-foreground">
                            <Upload className="mx-auto  h-12 w-12 mb-2" />
                            <p>No image selected</p>
                        </div>
                    )}
                </div>

                <div className="md:w-1/2 space-y-4">
                    <div className="cursor-pointer">
                        <Label
                            htmlFor="image-upload"
                            className="block text-sm font-medium mb-1"
                        >
                            Select an image
                        </Label>
                        <Input
                            ref={fileInputRef}
                            id="image-upload"
                            type="file"
                            accept="image/*"
                            onChange={handleFileChange}
                            className="cursor-pointer "
                            aria-describedby="file-upload-error"
                        />
                    </div>

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

                    {selectedImage && (
                        <p className="text-sm text-muted-foreground">
                            Image selected. You can choose another file to
                            replace it.
                        </p>
                    )}
                </div>
            </div>
            <div
                className={cn(
                    'relative h-2 w-full overflow-hidden rounded-full bg-primary/20 hidden',
                    { block: uploading }
                )}
            >
                <div
                    className={cn(
                        'h-full w-full flex-1 bg-primary transition-all -translate-x-[100%] duration-300',
                        { '-translate-x-[15%]': uploading }
                    )}
                ></div>
            </div>
        </div>
    );
}
