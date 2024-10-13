import { Res } from '@/types/http';
import { ResponseOrError } from './type';
import fetcher from '../http/fetcher';
import { getAuthH } from './helper';

type FileUploadResponse = {
    filename: string;
};

/**
 * Upload a image file to server and get filename
 *
 * @param key string
 * @param file File
 * @returns ResponseOrError<Res<FileUploadResponse>>
 */
export async function uploadImageFile(
    key: string,
    file: File
): Promise<ResponseOrError<Res<FileUploadResponse>>> {
    const formData = new FormData();
    formData.append('file', file);

    const [res, err] = await fetcher
        .set(...getAuthH(key))
        .post.formData<Res<FileUploadResponse>>('/upload', formData);

    if (res) {
        return [res, undefined];
    }
    return [undefined, err];
}
