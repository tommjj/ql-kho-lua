const ErrInternal = new Error('internal error');
const ErrDataNotFound = new Error('data not found');
const ErrNoUpdatedData = new Error('no data to update');
const ErrConflictingData = new Error(
    'data conflicts with existing data in unique column'
);
const ErrUnauthorized = new Error(
    'user is unauthorized to access the resource'
);

export {
    ErrInternal,
    ErrDataNotFound,
    ErrNoUpdatedData,
    ErrConflictingData,
    ErrUnauthorized,
};
